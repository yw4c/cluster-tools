IMAGE_NAME = "yw4code/cluster-tool"

upgrade:
	docker build -t "${IMAGE_NAME}:latest" .
	docker image push "${IMAGE_NAME}:latest"
