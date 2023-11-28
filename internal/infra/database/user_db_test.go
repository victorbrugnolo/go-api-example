package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/victorbrugnolo/go-api-example/internal/entity"
	"gorm.io/gorm"
	"testing"

	"gorm.io/driver/sqlite"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.User{})

	if err != nil {
		return
	}

	userDB := NewUser(db)

	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.User{})

	if err != nil {
		return
	}

	user, _ := entity.NewUser("John", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)

	fmt.Println(userFound)
	fmt.Println(user)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
