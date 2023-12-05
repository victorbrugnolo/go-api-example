package main

import (
	"github.com/victorbrugnolo/go-api-example/configs"
	"github.com/victorbrugnolo/go-api-example/internal/entity"
	"github.com/victorbrugnolo/go-api-example/internal/infra/database"
	"github.com/victorbrugnolo/go-api-example/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Product{}, &entity.User{})

	if err != nil {
		panic(err)
	}

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	err = http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}
