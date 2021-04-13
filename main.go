package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	r := gin.Default()
	r.GET("cluster-tool/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"xff":       c.Request.Header.Get("X-Forwarded-For"),
			"client IP": c.ClientIP(),
			"egress IP": getEgressIP(),
			"POD NAME":  os.Getenv("POD_NAME"),
			"NODE NAME": os.Getenv("NODE_NAME"),
		})
	})
	r.GET("cluster-tool/ping", ping)

	r.Run(":8081")
}

func getEgressIP() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.myip.com", nil)
	if err != nil {
		return ""
	}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return ""
	} else {
		return string(body)
	}
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket请求ping 返回pong
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
