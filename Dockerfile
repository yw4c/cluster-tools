######## Builder ########
FROM golang:1.17 as builder
WORKDIR /app
COPY / .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -ldflags="-X config.versionTag=$(git describe --tags --always)" -mod=vendor -v -o tool main.go
# get grpcurl bin
RUN go get -u github.com/fullstorydev/grpcurl/...
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
RUN cp $GOPATH/bin/grpcurl /app/grpcurl

######## Image ########
FROM alpine:3
RUN apk update && apk upgrade && \
    apk add --no-cache git bash curl mysql-client redis
COPY --from=builder /app/tool /app/tool
COPY --from=builder /app/env.yaml /app/env.yaml
COPY --from=builder /app/grpcurl /app/grpcurl