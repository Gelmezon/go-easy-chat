package main

import (
	"awesomeProject1/db"
	"awesomeProject1/domains"
	"awesomeProject1/ws"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var updater = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func test(c *gin.Context) {
	upgrade, err := updater.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	name := c.DefaultQuery("name", "")
	fmt.Println(name)
	// Close the subscription when we are done.
	defer func(upgrade *websocket.Conn) {
		err := upgrade.Close()
		if err != nil {
			return
		}
	}(upgrade)
	chanel := make(chan []byte, 100)
	ws.ChanelMap[name] = chanel
	go func() {
		c2 := ws.ChanelMap[name]
		if c2 == nil {
			for {
				c2 := ws.ChanelMap[name]
				if c2 != nil {
					break
				}
				time.Sleep(10 * time.Second)
			}
		}
		for bytes := range c2 {
			err2 := upgrade.WriteMessage(websocket.TextMessage, bytes)
			if err2 != nil {
				fmt.Println(err2)
			}
		}
	}()
	for {
		_, p, err := upgrade.ReadMessage()
		if err != nil {
			fmt.Println(err.Error())
		}
		message := &domains.Message{}
		errJson := json.Unmarshal(p, message)
		if errJson != nil {
			fmt.Println(errJson.Error())
		}
		ws.ChanelMap[message.Target] <- []byte(message.Content)
	}

}

func main() {

	db.RedisInit()
	defer func(RedisClient *redis.Client) {
		err := RedisClient.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(db.RedisClient)
	ws.ChanelMapInit()
	r := gin.Default()
	r.GET("/test", test)
	err := r.Run(":3333")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
