IMAGE_NAME = "yw4code/cluster-tool"

static:
	go fmt `go list ./... | grep -v ./vendor/...`
	go vet `go list ./... | grep -v ./vendor/...`

upgrade:
	docker build -t "${IMAGE_NAME}:latest" .
	docker image push "${IMAGE_NAME}:latest"

proto:
	protoc --go_out=plugins=grpc:. ./pb/observe.proto

