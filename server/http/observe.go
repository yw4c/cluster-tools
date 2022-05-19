package http

import (
	"cluster-tools/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Observe struct {
	ObserveService service.IObserveService
	upGrader       websocket.Upgrader
}

func NewObserveServer(service service.IObserveService) (server *Observe) {
	return &Observe{
		ObserveService: service,
		upGrader:  websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}


func (o *Observe) ObserveStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"xff":       c.Request.Header.Get("X-Forwarded-For"),
		"egress_ip": o.ObserveService.GetEgressIP(),
		"pod_name":  os.Getenv("POD_NAME"),
	})
}

func (o *Observe) ping(c *gin.Context) {

	// upgrade
	ws, err := o.upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		// Read received messages
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		// write response message
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
