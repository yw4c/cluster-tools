package app

import (
	"cluster-tools/pkg/di"
	httpServer "cluster-tools/server/http"
	"cluster-tools/service"
	"strconv"
)

func NewHttpApp(port int) IApp {
	return &HttpApp{
		port: port,
	}
}

type HttpApp struct {
	port int
	baseApp
}

func (s *HttpApp) registerSDKProvider() {
	s.baseApp.registerSDK()
}

func (s *HttpApp) registerServiceProvider() {
	s.baseApp.registerService()
}

func (s *HttpApp) registerServerProvider() {
	di.Providers = append(di.Providers, func(svc service.IObserveService) *httpServer.Observe {
		return httpServer.NewObserveServer(svc)
	})
}



func (s *HttpApp) Run(register TRegisterProviderFunc) {

	register(s)
	r := httpServer.NewRouter()
	r.Run(":" + strconv.Itoa(s.port))
}
