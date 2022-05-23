package sdk

import (
	"cluster-tools/config"
	"cluster-tools/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
	"sync"
)

type IUpstreamSDK interface {
	GetStatus(ctx context.Context, md metadata.MD)(*pb.GetStatusResponse, error)
}

var (
	upstreamOnce sync.Once
	upstreamInstance IUpstreamSDK
)


func NewUpstreamSDK() IUpstreamSDK {
	upstreamOnce.Do(func() {
		addr := strings.Join([]string{
			config.GetConfigInstance().UpstreamGRPC.Host,
			strconv.Itoa(config.GetConfigInstance().UpstreamGRPC.Port),
		}, ":")
		opts := grpc.WithTransportCredentials(insecure.NewCredentials())
		conn, err := grpc.Dial(addr, opts)
		if err != nil {
			panic(err.Error())
		}
		client := pb.NewObserveServiceClient(conn)
		upstreamInstance = &upstreamSDK{
			client: client,
		}
	})
	return upstreamInstance
}

type upstreamSDK struct {
	client pb.ObserveServiceClient
}

func (u *upstreamSDK) GetStatus(ctx context.Context, md metadata.MD) (*pb.GetStatusResponse, error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return u.client.GetStatus(ctx, &pb.GetStatusRequest{})
}