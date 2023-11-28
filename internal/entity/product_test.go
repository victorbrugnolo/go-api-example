package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Keyboard", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Keyboard", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Keyboard", 0)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Keyboard", -10)

	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProduct_Validate(t *testing.T) {
	p, err := NewProduct("Keyboard", 10)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
