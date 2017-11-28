FROM golang:alpine

MAINTAINER Dmitriy Detkov <maliciousgenius@gmail.com>

LABEL version="0.1" \
      description="web-stream-cv"

RUN apk --update --no-cache add git

ADD . $GOPATH/src

RUN go get github.com/lazywei/go-opencv && \
    go build -o $GOPATH/bin/main $GOPATH/src/main.go

EXPOSE 8000
ENTRYPOINT ["main"]
