package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

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

// swagger:route GET /products/{id} products getSingleProduct
// Return a single product from the database
// responses:
//
//	200: productResponse
//	404: errorResponse

// GetProduct returns a single product from the database
func (p *Products) GetProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	prod, err := data.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, prod)
}
