package services

import (
	"balabolka/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func RegisterWebSockets(router *gin.Engine) {
	router.GET("/ws", wsMsgHandler)
	router.GET("/", wsEchoHandler)
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[string]*websocket.Conn)

func wsMsgHandler(c *gin.Context) {
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("failed to set websocket upgrade: %+v", err)
		return
	}
	var name = conn.LocalAddr().String()
	clients[name] = conn
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
	delete(clients, name)
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
	clients[name] = conn
	println(fmt.Sprintf("New client handled, name = %s", name))
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		println(fmt.Sprintf("conn.ReadMessage %s complete", msg))
		err = conn.WriteMessage(t, msg)

		if err != nil {
			panic("failed to write message")
		}
	}
	delete(clients, name)
	println(fmt.Sprintf("Client  %s forgotten", name))

}
