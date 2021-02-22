######## Builder ########
FROM golang:1.14 as builder
WORKDIR /app
COPY / .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -ldflags="-X config.versionTag=$(git describe --tags)" -mod=vendor -v -o tool main.go


FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add --no-cache git bash curl mysql-client redis

# grpcurl
RUN go get github.com/fullstorydev/grpcurl/...
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

COPY --from=builder /app/tool /app/tool