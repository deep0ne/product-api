package handlers

import (
	"context"
	"log"
	"net/http"
	"product-api/data"

	"github.com/gin-gonic/gin"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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
