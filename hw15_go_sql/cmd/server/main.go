package main

import (
	"log"

	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/cmd/config"
	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/internal/handlers"
	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database, err := repository.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := repository.NewRepository(database)

	productHandler := handlers.NewProductHandler(repo)

	r := gin.Default()

	r.POST("/products", productHandler.ProductCreate)
	r.GET("/products/:id", productHandler.GetProduct)

	r.Run(":8080")
}
