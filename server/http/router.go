package http

import (
	"cluster-tools/pkg/di"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	var (
		observeServer *Observe
	)

	di.GetDIInstance().Invoke(func(
		observe *Observe,
	) {
		observeServer = observe
	})

	r := gin.Default()
	v1 := r.Group("observe")
	v1.GET("info", observeServer.ObserveStatus)
	v1.GET("ws-ping", observeServer.ping)
	//r.GET("ping", s.ping)
	return r
}
