FROM golang:alpine as builder

ARG GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux

WORKDIR /app

ARG ENABLE_PROXY=false

ADD go.* .
RUN if [ "$ENABLE_PROXY" = "true" ] ; then go env -w GOPROXY=https://goproxy.cn,direct ; fi \
    && go mod download
ADD . .
RUN go build -a -installsuffix cgo -o traefik-route-validate

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/traefik-route-validate .
ENTRYPOINT ["./traefik-route-validate"]