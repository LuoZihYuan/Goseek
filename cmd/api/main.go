package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/LuoZihYuan/Goseek/internal/handlers"
	"github.com/LuoZihYuan/Goseek/internal/repository"
	"github.com/LuoZihYuan/Goseek/internal/router"
	"github.com/LuoZihYuan/Goseek/internal/services"
)

const DataPath = "data/products.json"

// @title Products API
// @version 1.0.0
// @description API for managing products
// @contact.name API Support
// @contact.email zihyuan.luo@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @tag.name (Root)
// @tag.description Root / Server level services
// @tag.name Products
// @tag.description Product management operations
func main() {

	r := repository.NewProductsRepository(DataPath)
	s := services.NewProductService(r)
	rh := handlers.NewRootHandler()
	ph := handlers.NewProductsHandler(s)

	e := gin.Default()
	router.SetupRoutes(e, &router.AllHandlers{
		Root:     rh,
		Products: ph,
	})
	setupSwagger(e)

	log.Println("Starting server on :8080")
	if err := e.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
