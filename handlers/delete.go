package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//	204: noContent

// DeleteProduct deletes a product from database
func (p *Products) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := data.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusNoContent, "Product was successfully deleted")

}
