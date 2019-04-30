package main

import (
	"balabolka/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	println("HelloWorld")
	router := gin.Default()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	registerServices(router)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("GIN listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

func registerServices(r *gin.Engine) {
	services.RegisterInfoService(r)
	services.RegisterWebSockets(r)
}
