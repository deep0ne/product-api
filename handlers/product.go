// package classification of ProductAPI

// Documentation for Product API

// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package handlers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

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

func (p *Products) GetProducts(ctx *gin.Context) {
	lp := data.GetProducts()
	ctx.JSON(http.StatusOK, lp)
}

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

type KeyProduct struct{}

func (p *Products) MiddleWareProductValidations() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod := data.Product{}
		err := prod.FromJSON(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		err = prod.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		p.l.Println("Validation succeeded")
		ctx := context.WithValue(c.Request.Context(), KeyProduct{}, prod)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
