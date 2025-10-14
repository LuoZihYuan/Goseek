package router

import (
	"github.com/gin-gonic/gin"

	"github.com/LuoZihYuan/Goseek/internal/handlers"
)

type AllHandlers struct {
	Root     *handlers.RootHandler
	Products *handlers.ProductsHandler
}

func SetupRoutes(r *gin.Engine, h *AllHandlers) {
	root := r.Group("")
	{
		root.GET("/health", h.Root.GetHealth)
	}

	products := r.Group("/products")
	{
		products.GET("/search", h.Products.SearchProductsByQuery)
	}

}
