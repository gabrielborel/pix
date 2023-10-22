package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	BaseModel `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull" gorm:"column:owner_name;type:varchar(255);not null"`
	Bank      *Bank     `valid:"-" gorm:"ForeignKey:BankID"`
	BankID    string    `valid:"-" gorm:"column:bank_id;type:uuid;not null"`
	Number    string    `json:"number" valid:"notnull" gorm:"type:varchar(20)"`
	PixKeys   []*PixKey `valid:"-" gorm:"ForeignKey:AccountID"`
}

func NewAccount(ownerName, number string, bank *Bank) (*Account, error) {
	account := Account{
		OwnerName: ownerName,
		Bank:      bank,
		BankID:    bank.ID,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	if err := account.isValid(); err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *Account) isValid() error {
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}
	return nil
}
