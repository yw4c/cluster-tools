package grpc

import (
	"cluster-tools/c"
	"cluster-tools/pb"
	"cluster-tools/service"
	"context"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewObserveServer(svc service.IObserveService) (server *Observe) {

	return &Observe{
		observeService: svc,
	}
}

type Observe struct {
	observeService service.IObserveService
}

func (o Observe) GetStatus(ctx context.Context, request *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Internal, "get metadata fail")
	}
	log.Println(md.Get(strings.ToLower(strings.ToLower(c.XRequestID))))


	metadataHead := []string{
		c.XRequestID,
		c.XB3TraceID,
		c.XB3SpanID,
	}
	metadataMap := make(map[string]string)
	for _, v := range metadataHead {
		if ids := md.Get(strings.ToLower(v)); len(ids) > 0 {

			metadataMap[v] = ids[0]
		}
	}

	resp := &pb.GetStatusResponse{
		XRequestID: metadataMap[c.XRequestID],
		TraceID:    metadataMap[c.XB3TraceID],
		SpanID:     metadataMap[c.XB3SpanID],
		PodID:      os.Getenv("POD_NAME"),
	}
	log.Printf("%#v", resp)
	return resp, nil

}
