package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/victorbrugnolo/go-api-example/configs"
	"github.com/victorbrugnolo/go-api-example/internal/entity"
	"github.com/victorbrugnolo/go-api-example/internal/infra/database"
	"github.com/victorbrugnolo/go-api-example/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	cfgs, err := configs.LoadConfig(".")

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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, cfgs.TokenAuth, cfgs.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfgs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJwt)

	err = http.ListenAndServe(":8000", r)

	if err != nil {
		panic(err)
	}
}
