package grpc

import (
	"cluster-tools/c"
	"cluster-tools/pb"
	"cluster-tools/pkg/errors"
	"cluster-tools/service"
	"context"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc/metadata"
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
		return nil, errors.Wrap(errors.ErrInternalError, "load header fail")
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
		PodName:    os.Getenv("POD_NAME"),
	}
	return resp, nil

}
