package repository

import (
	"fmt"

	"github.com/gabrielborel/pix/codepix/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func NewTransactionRepositoryDb(db *gorm.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{Db: db}
}

func (r TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
