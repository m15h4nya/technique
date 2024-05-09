FROM golang:1.21.0-alpine as build

COPY . /vault
RUN cd /vault && go build -o app .

FROM debian:buster-slim

RUN mkdir -p /opt/vault
COPY --from=build /vault/app /opt/vault/app

WORKDIR /opt/vault
CMD ["./app"]