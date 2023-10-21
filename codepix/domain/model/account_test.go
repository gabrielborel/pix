package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidAccount(t *testing.T) {
	ownerName := "Test Owner"
	bank, _ := NewBank("001", "Test Bank")
	account, err := NewAccount(ownerName, "0001", bank)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, ownerName, account.OwnerName)
	assert.Equal(t, bank, account.Bank)
	assert.Equal(t, "0001", account.Number)
}

func TestCreateAccountWithouOwnerName(t *testing.T) {
	ownerName := ""
	bank, _ := NewBank("001", "Test Bank")
	account, err := NewAccount(ownerName, "0001", bank)

	assert.NotNil(t, err)
	assert.Equal(t, "owner_name: Missing required field", err.Error())
	assert.Nil(t, account)
}

func TestCreateAccountWithoutNumber(t *testing.T) {
	ownerName := "Test Owner"
	bank, _ := NewBank("001", "Test Bank")
	account, err := NewAccount(ownerName, "", bank)

	assert.NotNil(t, err)
	assert.Equal(t, "number: Missing required field", err.Error())
	assert.Nil(t, account)
}
