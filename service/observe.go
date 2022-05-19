package service

import (
	"cluster-tools/sdk"
)

type IObserveService interface {
	GetEgressIP() string
}

func NewObserveService(sdkUpstream sdk.IUpstreamSDK, sdkEgress sdk.IEgressAddress) IObserveService {
	return &observeService{
		sdkUpstream: sdkUpstream,
		sdkEgressAddress: sdkEgress,
	}
}

type observeService struct {
	sdkUpstream sdk.IUpstreamSDK
	sdkEgressAddress sdk.IEgressAddress
}

func (s *observeService) GetEgressIP() string {
	return s.sdkEgressAddress.Get()
}
