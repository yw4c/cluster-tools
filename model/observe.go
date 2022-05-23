package model

import "cluster-tools/pb"

type ObserveStatusResponse struct {
	Xff           string                `json:"xff"`
	EgressAddress string                `json:"egressAddress"`
	PodName       string                `json:"podName"`
	Upstream      *pb.GetStatusResponse `json:"upstream"`
}
