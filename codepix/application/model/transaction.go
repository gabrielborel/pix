package application

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"account_id"`
	Amount       float64 `json:"amount" validate:"required"`
	PixKeyTo     string  `json:"pix_key_to" validate:"required"`
	PixKeyKindTo string  `json:"pix_key_kind_to" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error"`
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		return fmt.Errorf("error during the Transaction validation: %s", err.Error())
	}
	return nil
}

func (t *Transaction) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	err = t.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)
	if err != nil {
		return nil, nil
	}

	return result, nil
}
