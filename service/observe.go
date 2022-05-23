package service

import (
	"cluster-tools/pb"
	pkgGrpc "cluster-tools/pkg/grpc"
	"cluster-tools/sdk"
	"context"
	"net/http"
)

type IObserveService interface {
	GetEgressIP() string
	GetUpstreamInfo(ctx context.Context, req *http.Request) (*pb.GetStatusResponse, error)
}

func NewObserveService(sdkUpstream sdk.IUpstreamSDK, sdkEgress sdk.IEgressAddress) IObserveService {
	return &observeService{
		sdkUpstream:      sdkUpstream,
		sdkEgressAddress: sdkEgress,
	}
}

type observeService struct {
	sdkUpstream      sdk.IUpstreamSDK
	sdkEgressAddress sdk.IEgressAddress
}

func (s *observeService) GetUpstreamInfo(ctx context.Context, req *http.Request) (*pb.GetStatusResponse, error) {
	md := pkgGrpc.InjectHeadersIntoMetadata(ctx, req)
	return s.sdkUpstream.GetStatus(ctx, md)
}

func (s *observeService) GetEgressIP() string {
	return s.sdkEgressAddress.Get()
}
