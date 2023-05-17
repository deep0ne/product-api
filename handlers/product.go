package handlers

import (
	"context"
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
	prod := ctx.Request.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
	ctx.Status(http.StatusCreated)
}

func (p *Products) GetProducts(ctx *gin.Context) {
	lp := data.GetProducts()
	ctx.JSON(http.StatusOK, lp)
}

func (p *Products) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindUri(&id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	prod := ctx.Request.Context().Value(KeyProduct{}).(data.Product)
	err := data.UpdateProduct(id, &prod)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusOK)
}

type KeyProduct struct{}

func (p *Products) MiddleWareProductValidations() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod := data.Product{}
		err := prod.FromJSON(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		p.l.Println("Validation succeeded")
		ctx := context.WithValue(c.Request.Context(), KeyProduct{}, prod)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
