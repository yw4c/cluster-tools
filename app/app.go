package app

import (
	"cluster-tools/pkg/di"
	"cluster-tools/sdk"
	"cluster-tools/service"
)

type IApp interface {
	registerSDKProvider()
	registerServiceProvider()
	registerServerProvider()
	Run(TRegisterProviderFunc)
}

type TRegisterProviderFunc func (app IApp)

var RegisterProvidersFunc = func(app IApp) {
	app.registerSDKProvider()
	app.registerServiceProvider()
	app.registerServerProvider()
}

type baseApp struct{}

func (*baseApp) registerSDK() {
	di.Providers = append(di.Providers, func() sdk.IUpstreamSDK {
		return sdk.NewUpstreamSDK()
	})
	di.Providers = append(di.Providers, func() sdk.IEgressAddress {
		return sdk.NewEgressAddress()
	})
}

func (*baseApp) registerService() {
	di.Providers = append(di.Providers, func(
		sdkUpstream sdk.IUpstreamSDK,
		sdkEgress sdk.IEgressAddress,
	) service.IObserveService {
		return service.NewObserveService(sdkUpstream, sdkEgress)
	})
}

