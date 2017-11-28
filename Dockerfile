FROM golang:alpine

MAINTAINER Dmitriy Detkov <maliciousgenius@gmail.com>

LABEL version="0.1" \
      description="web-stream-cv"

RUN apk --update --no-cache add git

RUN echo $PATH

RUN mkdir -p /opt/app
ADD . /opt/app/
WORKDIR /opt/app
RUN go get && \
    go generate && \
    go build -o main .

EXPOSE 8000
ENTRYPOINT ["/opt/app/main"]
