FROM golang:1.13-alpine3.11
WORKDIR /go/src/app
COPY . .
RUN mv ./config/sample.conf.yaml ./config/conf.yaml
CMD go run .
