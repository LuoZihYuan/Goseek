package router

import (
	"github.com/gin-gonic/gin"

	"github.com/LuoZihYuan/Goseek/internal/handlers"
)

func SetupRoutes(r *gin.Engine, h *handlers.ProductsHandler) {
	products := r.Group("/products")
	products.GET("/search", h.SearchByQuery)
}
