package grpc

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"strings"
)

func InjectHeadersIntoMetadata(ctx context.Context, req *http.Request) metadata.MD {
	//https://aspenmesh.io/2018/04/tracing-grpc-with-istio/
	var (
		otHeaders = []string{
			"x-request-id",
			"x-b3-traceid",
			"x-b3-spanid",
			"x-b3-parentspanid",
			"x-b3-sampled",
			"x-b3-flags",
			"x-ot-span-context"}
	)
	var pairs []string

	for k := range req.Header {
		for _, h := range otHeaders {
			if strings.ToLower(k) == h {
				//logrus.Debug("merging otHeader:" , h)
				v := req.Header.Get(k)
				pairs = append(pairs, h, v)
			}
		}
	}
	log.Println(pairs)

	md := metadata.Pairs(pairs...)

	return md
}
