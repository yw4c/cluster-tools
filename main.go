package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("cluster-tool/info", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"xff":       c.Request.Header.Get("X-Forwarded-For"),
			"client IP": c.ClientIP(),
			"egress IP": getEgressIP(),
			"POD NAME":  os.Getenv("POD_NAME"),
		})
	})

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
