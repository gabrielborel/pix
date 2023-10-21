package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	BaseModel `valid:"required"`
	Code      string     `json:"code" valid:"notnull"`
	Name      string     `json:"name" valid:"notnull"`
	Accounts  []*Account `valid:"-"`
}

func NewBank(code, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	if err := bank.isValid(); err != nil {
		return nil, err
	}

	return &bank, nil
}

func (b *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(b)
	if err != nil {
		return err
	}
	return nil
}
