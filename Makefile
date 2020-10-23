IMAGE_NAME = "gcr.io/silkrode-golang/cluster-tools"
NAMESPACE ?= "orchestrator"
upgrade:
	docker build -t "${IMAGE_NAME}:latest" .
	docker image push "${IMAGE_NAME}:latest"

deploy:
	kubectl apply -f ./deployment  -n ${NAMESPACE}