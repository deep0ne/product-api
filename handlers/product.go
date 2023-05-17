package handlers

import (
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
	prod := data.Product{}
	err := prod.FromJSON(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
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

	prod := data.Product{}
	err := prod.FromJSON(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = data.UpdateProduct(id, &prod)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusOK)
}
