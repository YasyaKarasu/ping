FROM golang:1.18 AS builder

ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on

RUN mkdir -p /tmp/app

WORKDIR /tmp/app

COPY . .

RUN go mod download all && go build -tags netgo -o /pingpong

FROM alpine:3.13.1

COPY --from=builder /pingpong /pingpong

ENTRYPOINT ["/pingpong"]