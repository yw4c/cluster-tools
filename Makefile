IMAGE_NAME = "yw4code/cluster-tool"

upgrade:
	go fmt `go list ./... | grep -v ./vendor/...`
	go vet `go list ./... | grep -v ./vendor/...`
	docker build -t "${IMAGE_NAME}:latest" .
	docker image push "${IMAGE_NAME}:latest"

proto:
	protoc --go_out=plugins=grpc:. ./pb/observe.proto

