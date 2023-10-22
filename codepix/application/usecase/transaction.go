package usecase

import (
	"errors"
	"log"

	"github.com/gabrielborel/pix/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixKeyRepository      model.PixKeyRepositoryInterface
}

func NewTransactionUseCase(
	transactionRepository model.TransactionRepositoryInterface, 
	pixKeyRepository model.PixKeyRepositoryInterface,
) *TransactionUseCase {
	return &TransactionUseCase{
		TransactionRepository: transactionRepository,
		PixKeyRepository:      pixKeyRepository,
	}
}

func (uc *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := uc.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := uc.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	uc.TransactionRepository.Save(transaction)
	if transaction.ID == "" {
		return nil, errors.New("unable to process register this transaction at the moment")
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	transaction, err := uc.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err = uc.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	transaction, err := uc.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err = uc.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (uc *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	transaction, err := uc.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason
	err = uc.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
