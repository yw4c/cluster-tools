FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add --no-cache git bash curl mysql-client redis

# grpcurl
RUN go get github.com/fullstorydev/grpcurl/...
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl

CMD ["sleep", "100000s"]