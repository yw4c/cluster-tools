######## Builder ########
FROM golang:1.14 as builder
WORKDIR /app
COPY / .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -ldflags="-X config.versionTag=$(git describe --tags)" -mod=vendor -v -o tool main.go
# grpcurl
RUN cd $GOPATH && ls -al
RUN go get github.com/fullstorydev/grpcurl/...
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest -mod=mod
COPY $GOPATH/bin/grpcurl /app/grpcurl

######## Image ########
FROM alpine:3
RUN apk update && apk upgrade && \
    apk add --no-cache git bash curl mysql-client redis
COPY --from=builder /app/tool /app/tool
COPY --from=builder /app/grpcurl /app/grpcurl