package main

import (
	"balabolka/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
