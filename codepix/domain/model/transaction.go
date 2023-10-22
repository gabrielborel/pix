package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending    string = "pending"
	TransactionCompleted  string = "completed"
	TransactionError      string = "error"
	TransactionConfirmed  string = "confirmed"
	TransactionCancelled  string = "cancelled"
	TransactionProcessing string = "processing"
)

type Transactions struct {
	Transaction []*Transaction
}

type Transaction struct {
	BaseModel         `valid:"required"`
	AccountFrom       *Account `valid:"-" gorm:"ForeignKey:AccountFromID"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" valid:"notnull" gorm:"type:float"`
	PixKeyTo          *PixKey  `valid:"-" gorm:"ForeignKey:PixKeyToID"`
	PixKeyToID        string   `gorm:"column:pix_key_to_id;type:uuid;" valid:"notnull"`
	Status            string   `json:"status" valid:"notnull" gorm:"type:varchar(20)"`
	Description       string   `json:"description" valid:"notnull" gorm:"type:varchar(255)"`
	CancelDescription string   `json:"cancel_description" valid:"-" gorm:"type:varchar(255)"`
}

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom:   accountFrom,
		AccountFromID: accountFrom.ID,
		Amount:        amount,
		PixKeyTo:      pixKeyTo,
		PixKeyToID:    pixKeyTo.ID,
		Status:        TransactionPending,
		Description:   description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	if err := transaction.isValid(); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Confirm() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionCancelled
	t.UpdatedAt = time.Now()
	t.CancelDescription = description
	err := t.isValid()
	return err
}
