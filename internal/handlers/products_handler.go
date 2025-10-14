package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/LuoZihYuan/Goseek/internal/models"
	"github.com/LuoZihYuan/Goseek/internal/services"
	"github.com/LuoZihYuan/Goseek/internal/shared"
)

type ProductsHandler struct {
	service *services.ProductsService
}

func NewProductsHandler(service *services.ProductsService) *ProductsHandler {
	return &ProductsHandler{service: service}
}

// SearchProductsByQuery handles GET /products/search?q=
// @Summary Search products by query
// @Description Search for products by name or category using a query string
// @ID searchProducts
// @Tags Products
// @Produce json
// @Param q query string true "Search query for product name or category"
// @Success 200 {array} models.Product
// @Failure 400 {object} models.HttpError
// @Failure 404 {object} models.HttpError
// @Failure 500 {object} models.HttpError
// @Router /products/search [get]
func (h *ProductsHandler) SearchProductsByQuery(c *gin.Context) {
	const maxSearches = 100
	const maxResults = 20

	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, models.HttpError{
			Error:   400,
			Message: "BAD_REQUEST",
			Details: "Query is required.",
		})

		return
	}

	startTime := time.Now()
	products, err := h.service.SearchProductsByQuery(query, maxSearches, maxResults)
	searchTime := time.Since(startTime)

	if err == shared.ErrProductNotFound {
		c.JSON(http.StatusNotFound, models.HttpError{
			Error:   404,
			Message: "NOT_FOUND",
			Details: fmt.Sprintf("`%v` does not exist in any product names and categories.", query),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, models.HttpError{
			Error:   500,
			Message: "INTERNAL_SERVER_ERROR",
			Details: fmt.Sprintf("Unhandled error: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, models.ProductsSearchResponse{
		Products:   products,
		TotalFound: len(products),
		SearchTime: fmt.Sprintf("%.2fms", searchTime.Seconds()*1_000),
	})
}
