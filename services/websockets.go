package services

import (
	"balabolka/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

func RegisterWebSockets(router *gin.Engine) {
	router.GET("/ws", wsMsgHandler)
	router.GET("/", wsEchoHandler)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients sync.Map

func wsMsgHandler(c *gin.Context) {
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("failed to set websocket upgrade: %+v", err)
		return
	}
	var name = conn.LocalAddr().String()
	clients.Store(name, conn)

	println(fmt.Sprintf("New client handled, name = %s", name))
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var myMessage models.Message
		err = json.Unmarshal(msg, &myMessage)
		if err != nil {
			panic("failed unmarshal json")
		}

		bt, err := json.Marshal(myMessage)

		err = conn.WriteMessage(t, bt)

		if err != nil {
			panic("failed to write message")
		}
	}
	clients.Delete(name)
	println(fmt.Sprintf("Client  %s forgotten", name))

}

func wsEchoHandler(c *gin.Context) {
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }
	println("wsEchoHandler start")
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("failed to set websocket upgrade: %+v", err)
		return
	}
	var name = conn.LocalAddr().String()
	clients.Store(name, conn)
	println(fmt.Sprintf("New client handled, name = %s", name))
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		println(fmt.Sprintf("conn.ReadMessage %s complete", msg))

		clients.Range(func(key, value interface{}) bool {
			conn = value.(*websocket.Conn)
			err = conn.WriteMessage(t, msg)
			println(fmt.Sprintf("sending %s to %s", msg, conn.LocalAddr().String()))
			if err != nil {
				return false
			}
			return true
		})

	}
	clients.Delete(name)
	println(fmt.Sprintf("Client  %s forgotten", name))

}
