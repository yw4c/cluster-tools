package http

import (
	"cluster-tools/model"
	"cluster-tools/service"
	"log"
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

	upstream, err := o.ObserveService.GetUpstreamInfo(c, c.Request)
	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(200, &model.ObserveStatusResponse{
		Xff:           c.Request.Header.Get("X-Forwarded-For"),
		EgressAddress: o.ObserveService.GetEgressIP(),
		PodName:       os.Getenv("POD_NAME"),
		Upstream:      upstream,
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
