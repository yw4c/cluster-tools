package app

import (
	"cluster-tools/pb"
	"cluster-tools/pkg/di"
	grpcServer "cluster-tools/server/grpc"
	"cluster-tools/service"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcApp(port int) IApp {
	return &GrpcApp{
		port: port,
	}
}

type GrpcApp struct {
	port int
	baseApp
}

func (s *GrpcApp) registerSDKProvider() {
	s.baseApp.registerSDK()
}

func (s *GrpcApp) registerServiceProvider() {
	s.baseApp.registerService()
}

func (s *GrpcApp) registerServerProvider() {

	di.Providers = append(di.Providers, func(svc service.IObserveService) *grpcServer.Observe {
		return grpcServer.NewObserveServer(svc)
	})
}

func (s *GrpcApp) Run(register TRegisterProviderFunc) {

	var (
		observeServer *grpcServer.Observe
	)

	register(s)

	di.GetDIInstance().Invoke(func(
		observe *grpcServer.Observe,
	) {
		observeServer = &grpcServer.Observe{}
	})

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcSvc := grpc.NewServer()
	pb.RegisterObserveServiceServer(grpcSvc, observeServer)
	reflection.Register(grpcSvc)

	go func() {
		if err := grpcSvc.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
