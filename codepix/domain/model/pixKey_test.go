package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidPixKey(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	account, _ := NewAccount("Test Owner", "0001", bank)
	kind := "cpf"
	key := "11122233344"
	pixKey, err := NewPixKey(kind, key, account)

	assert.Nil(t, err)
	assert.NotNil(t, pixKey)
	assert.Equal(t, pixKey.Account, account)
	assert.Equal(t, pixKey.AccountID, account.ID)
	assert.Equal(t, pixKey.Kind, kind)
	assert.Equal(t, pixKey.Key, key)
	assert.Equal(t, pixKey.Status, "active")
}

func TestCreatePixAccountWithoutKind(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	account, _ := NewAccount("Test Owner", "0001", bank)
	kind := ""
	key := "11122233344"
	pixKey, err := NewPixKey(kind, key, account)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid type of key", err.Error())
	assert.Nil(t, pixKey)
}

func TestCreatePixAccountWithoutKey(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	account, _ := NewAccount("Test Owner", "0001", bank)
	kind := "cpf"
	key := ""
	pixKey, err := NewPixKey(kind, key, account)

	assert.NotNil(t, err)
	assert.Equal(t, "key: Missing required field", err.Error())
	assert.Nil(t, pixKey)
}

func TestCreatePixAccountWithInvalidKind(t *testing.T) {
	bank, _ := NewBank("001", "Test Bank")
	account, _ := NewAccount("Test Owner", "0001", bank)
	kind := "invalid"
	key := "11122233344"
	pixKey, err := NewPixKey(kind, key, account)

	assert.NotNil(t, err)
	assert.Equal(t, "invalid type of key", err.Error())
	assert.Nil(t, pixKey)
}
