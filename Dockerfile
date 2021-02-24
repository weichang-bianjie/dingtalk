FROM golang:1.15.5-alpine3.12 as builder

# Set up dependencies
ENV PACKAGES make git libc-dev bash gcc
ENV GO111MODULE  on

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
RUN apk add --no-cache $PACKAGES && make all

FROM alpine:3.12
ENV TZ           Asia/Shanghai
WORKDIR /root/
COPY --from=builder /go/src/dingtalk /usr/local/bin
CMD dingtalk