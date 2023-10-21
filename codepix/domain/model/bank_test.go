package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateValidBank(t *testing.T) {
	code := "001"
	name := "Test Bank"
	bank, err := NewBank(code, name)

	assert.Nil(t, err)
	assert.NotNil(t, bank)
	assert.Equal(t, code, bank.Code)
	assert.Equal(t, name, bank.Name)
}

func TestCreateBankWithouCode(t *testing.T) {
	code := ""
	name := "Test Bank"
	bank, err := NewBank(code, name)

	assert.NotNil(t, err)
	assert.Equal(t, "code: Missing required field", err.Error())
	assert.Nil(t, bank)
}

func TestCreateBankWithouName(t *testing.T) {
	code := "001"
	name := ""
	bank, err := NewBank(code, name)

	assert.NotNil(t, err)
	assert.Equal(t, "name: Missing required field", err.Error())
	assert.Nil(t, bank)
}
