package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/victorbrugnolo/go-api-example/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestProduct_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return
	}

	product, err := entity.NewProduct("Mouse", 10.0)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestProduct_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return
	}

	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)
}

func TestProduct_FindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return
	}

	product, err := entity.NewProduct("Mouse", 10.0)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Mouse", product.Name)
}

func TestProduct_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return
	}

	product, err := entity.NewProduct("Mouse", 10.0)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product.Name = "Mouse Gamer"
	err = productDB.Update(product)
	assert.NoError(t, err)

	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Mouse Gamer", product.Name)
}

func TestProduct_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return
	}

	product, err := entity.NewProduct("Mouse", 10.0)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product)
	assert.NoError(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
