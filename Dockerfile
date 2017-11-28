FROM golang

MAINTAINER Dmitriy Detkov <maliciousgenius@gmail.com>

LABEL version="0.1" \
      description="web-stream-cv"

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y --no-install-recommends libopencv-dev \
    pkg-config

ADD . $GOPATH/src

RUN go get github.com/lazywei/go-opencv && \
    go build -o $GOPATH/bin/main $GOPATH/src/main.go

EXPOSE 8000
ENTRYPOINT ["main"]
