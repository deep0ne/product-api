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

// func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
// 	prod := data.Product{}
// 	err := prod.FromJSON(r.Body)
// 	if err != nil {
// 		http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
// 	}
// 	p.l.Printf("Prod: %#v", prod)
// 	data.AddProduct(&prod)
// }

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
	}

	err = data.UpdateProduct(id, &prod)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusOK, prod)
}
