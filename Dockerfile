FROM golang:alpine as builder

ARG GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOPROXY=https://goproxy.cn

WORKDIR /app

ADD go.* .
RUN go mod download
ADD . .
RUN go build -a -installsuffix cgo -o traefik-route-validate

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/traefik-route-validate .
ENTRYPOINT ["./traefik-route-validate"]