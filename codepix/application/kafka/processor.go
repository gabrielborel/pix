package kafka

import (
	"fmt"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gabrielborel/pix/codepix/application/factory"
	appmodel "github.com/gabrielborel/pix/codepix/application/model"
	"github.com/gabrielborel/pix/codepix/application/usecase"
	"github.com/gabrielborel/pix/codepix/domain/model"
	"github.com/jinzhu/gorm"
)

type KafkaProcessor struct {
	Database     *gorm.DB
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(database *gorm.DB, producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:     database,
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{
		os.Getenv("kafkaTransactionTopic"),
		os.Getenv("kafkaTransactionConfirmationTopic"),
	}
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.processMessage(msg)
		} else {
			fmt.Println(err.Error())
		}
	}
}

func (k *KafkaProcessor) processMessage(msg *ckafka.Message) error {
	transactionsTopic := os.Getenv("kafkaTransactionTopic")
	transactionConfirmationToic := os.Getenv("kafkaTransactionConfirmationTopic")

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopic:
		return k.processTransaction(msg)
	case transactionConfirmationToic:
		return k.processTransactionConfirmation(msg)
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}

	return nil
}

func (k *KafkaProcessor) processTransaction(msg *ckafka.Message) error {
	transaction := appmodel.NewTransaction()
	err := transaction.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(k.Database)

	createdTransaction, err := transactionUseCase.Register(
		transaction.AccountID,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyKindTo,
		transaction.Description,
	)
	if err != nil {
		fmt.Println("error registering transaction", err)
		return err
	}

	bankDestinationTopic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code
	transaction.ID = createdTransaction.ID
	transaction.Status = model.TransactionPending
	transactionJson, err := transaction.ToJson()
	if err != nil {
		fmt.Println("error converting transaction to json", err)
		return err
	}

	err = Publish(string(transactionJson), bankDestinationTopic, k.Producer, k.DeliveryChan)
	if err != nil {
		fmt.Println("error publishing transaction to kafka", err)
		return err
	}

	return nil
}

func (k *KafkaProcessor) processTransactionConfirmation(msg *ckafka.Message) error {
	transaction := appmodel.NewTransaction()
	err := transaction.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	transactionUseCase := factory.TransactionUseCaseFactory(k.Database)

	if transaction.Status == model.TransactionConfirmed {
		err = k.confirmTransaction(transaction, transactionUseCase)
		if err != nil {
			return err
		}
	} else if transaction.Status == model.TransactionCompleted {
		_, err = transactionUseCase.Complete(transaction.ID)
		if err != nil {
			return err
		}
		return nil
	} else if transaction.Status == model.TransactionError {
		_, err = transactionUseCase.Error(transaction.ID, transaction.Error)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (k *KafkaProcessor) confirmTransaction(
	transaction *appmodel.Transaction,
	transactionUseCase *usecase.TransactionUseCase,
) error {
	confirmedTransaction, err := transactionUseCase.Confirm(transaction.ID)
	if err != nil {
		return err
	}

	topic := "bank" + confirmedTransaction.AccountFrom.Bank.Code
	transactionJson, err := transaction.ToJson()
	if err != nil {
		return err
	}

	err = Publish(string(transactionJson), topic, k.Producer, k.DeliveryChan)
	if err != nil {
		return err
	}

	return nil
}
