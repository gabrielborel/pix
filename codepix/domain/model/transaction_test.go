package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidTransaction(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	fromAccount, _ := NewAccount("Test Owner", "1111", bank)
	amount := 3.10
	toAccount, _ := NewAccount("Test Owner", "2222", bank)
	pixKey, _ := NewPixKey("cpf", "11122233344", toAccount)
	transaction, err := NewTransaction(fromAccount, amount, pixKey, "Test Description")

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, transaction.AccountFrom.ID, fromAccount.ID)
	assert.Equal(t, transaction.Amount, amount)
	assert.Equal(t, transaction.PixKeyTo.ID, pixKey.ID)
	assert.Equal(t, transaction.PixKeyTo.Account.ID, toAccount.ID)
	assert.Equal(t, transaction.Status, TransactionPending)
	assert.Equal(t, transaction.Description, "Test Description")
	assert.Empty(t, transaction.CancelDescription)
}

func TestCreateTransactionWithoutAmount(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	fromAccount, _ := NewAccount("Test Owner", "1111", bank)
	toAccount, _ := NewAccount("Test Owner", "2222", bank)
	pixKey, _ := NewPixKey("cpf", "11122233344", toAccount)
	_, err := NewTransaction(fromAccount, 0, pixKey, "Test Description")

	assert.NotNil(t, err)
	assert.Equal(t, "the amount must be greater than 0", err.Error())
}

func TestCreateTransactionWithTheSameSourceAndDestinationAccount(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	account, _ := NewAccount("Test Owner", "1111", bank)
	pixKey, _ := NewPixKey("cpf", "11122233344", account)
	_, err := NewTransaction(account, 100, pixKey, "Test Description")

	assert.NotNil(t, err)
	assert.Equal(t, "the source and destination account cannot be the same", err.Error())
}

func TestCompleteTransaction(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	fromAccount, _ := NewAccount("Test Owner", "1111", bank)
	amount := 3.10
	toAccount, _ := NewAccount("Test Owner", "2222", bank)
	pixKey, _ := NewPixKey("cpf", "11122233344", toAccount)
	transaction, err := NewTransaction(fromAccount, amount, pixKey, "Test Description")
	transaction.Complete()

	assert.Nil(t, err)
	assert.Equal(t, transaction.Status, TransactionCompleted)
	assert.NotNil(t, transaction.UpdatedAt)
}

func TestConfirmTransaction(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	fromAccount, _ := NewAccount("Test Owner", "1111", bank)
	amount := 3.10
	toAccount, _ := NewAccount("Test Owner", "2222", bank)
	pixKey, _ := NewPixKey("cpf", "11122233344", toAccount)
	transaction, err := NewTransaction(fromAccount, amount, pixKey, "Test Description")
	transaction.Complete()
	transaction.Confirm()

	assert.Nil(t, err)
	assert.Equal(t, transaction.Status, TransactionConfirmed)
	assert.NotNil(t, transaction.UpdatedAt)
}

func TestCancelTransaction(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	fromAccount, _ := NewAccount("Test Owner", "1111", bank)
	amount := 3.10
	toAccount, _ := NewAccount("Test Owner", "2222", bank)
	pixKey, _ := NewPixKey("cpf", "11122233344", toAccount)
	transaction, err := NewTransaction(fromAccount, amount, pixKey, "Test Description")
	transaction.Cancel("Test Cancel Description")

	assert.Nil(t, err)
	assert.Equal(t, transaction.Status, TransactionCancelled)
	assert.NotNil(t, transaction.UpdatedAt)
	assert.Equal(t, transaction.CancelDescription, "Test Cancel Description")
}
