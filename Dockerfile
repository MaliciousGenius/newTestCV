FROM golang:alpine

MAINTAINER Dmitriy Detkov <maliciousgenius@gmail.com>

LABEL version="0.1" \
      description="web-stream-cv"

# OpenCV
ENV OPENCV_VERSION 2.4.13.4
RUN apk --update --no-cache add linux-headers gcc g++ make cmake  wget unzip
RUN mkdir /tmp/opencv \
    && cd /tmp/opencv \
    && wget -q https://github.com/opencv/opencv/archive/$OPENCV_VERSION.zip \
    && unzip $OPENCV_VERSION.zip \
    && rm $OPENCV_VERSION.zip \
    && mkdir build \
    && cd build \
    && cmake ../opencv-$OPENCV_VERSION \
    && make -j2 \
    && make install \
    && cd / \
    && rm -rf /tmp/opencv

ADD . $GOPATH/src

RUN apk --update --no-cache add git
#build-base

RUN go get github.com/lazywei/go-opencv && \
    go build -o $GOPATH/bin/main $GOPATH/src/main.go

EXPOSE 8000
ENTRYPOINT ["main"]
