package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(l)

	sm := gin.Default()

	sm.GET("/products", ph.GetProducts)
	sm.DELETE("/products/:id", ph.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "./swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	sm.GET("/docs", func(ctx *gin.Context) {
		sh.ServeHTTP(ctx.Writer, ctx.Request)
	})
	sm.GET("/swagger.yaml", func(ctx *gin.Context) {
		http.FileServer(http.Dir("./")).ServeHTTP(ctx.Writer, ctx.Request)
	})

	changeRoutes := sm.Group("/products")
	changeRoutes.Use(ph.MiddleWareProductValidations())
	changeRoutes.PUT("/:id", ph.UpdateProduct)
	changeRoutes.POST("", ph.CreateProduct)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGTERM, os.Interrupt)

	sig := <-sigchan
	l.Println("Recieved terminate signal, graceful shutdown...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
