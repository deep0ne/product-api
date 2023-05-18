package handlers

import (
	"errors"
	"net/http"
	"product-api/data"

	"github.com/gin-gonic/gin"
)

func (p *Products) CreateProduct(ctx *gin.Context) {
	prod, ok := ctx.Request.Context().Value(KeyProduct{}).(data.Product)
	if !ok {
		ctx.JSON(http.StatusBadRequest, errors.New("invalid data input"))
		return
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
	ctx.JSON(http.StatusCreated, "Product was created successfully")
}
