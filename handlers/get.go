package handlers

import (
	"net/http"
	"product-api/data"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(ctx *gin.Context) {
	lp := data.GetProducts()
	ctx.JSON(http.StatusOK, lp)
}
