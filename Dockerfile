FROM golang:1.21.0-alpine as build

COPY . /vault
WORKDIR /vault

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init -g server.go

RUN go build -o app .

FROM debian:buster-slim

RUN mkdir -p /opt/vault
COPY --from=build /vault/app /opt/vault/app
COPY --from=build /vault/vault /opt/vault/vault

WORKDIR /opt/vault
CMD ["./app"]