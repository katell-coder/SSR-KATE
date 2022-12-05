package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	engine := gin.Default()

	engine.GET("/helloWebSocket", func(context *gin.Context) {
		// 将普通的http GET请求升级为websocket请求
		client, _ := upgrader.Upgrade(context.Writer, context.Request, nil)
		for {
			// 每隔两秒给前端推送一句消息“hello, WebSocket”
			err := client.WriteMessage(websocket.TextMessage, []byte("hello, WebSocket"))
			if err != nil {
				log.Println(err)
			}
			time.Sleep(time.Second * 2)
		}
	})
	err := engine.Run(":8090")
	if err != nil {
		log.Fatalln(err)
	}
}
