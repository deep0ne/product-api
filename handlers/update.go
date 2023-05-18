package handlers

import (
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (p *Products) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	prod := ctx.Request.Context().Value(KeyProduct{}).(data.Product)
	err := data.UpdateProduct(id, &prod)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Product was updated successfully")
}
