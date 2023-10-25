# syntax=docker.io/docker/dockerfile:1
FROM golang:1.21 AS builder
WORKDIR /go/src/app
COPY ./app .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -v -o app .

FROM scratch

WORKDIR /
COPY --from=builder /go/src/app/app /app

EXPOSE 8080
ENTRYPOINT [ "/app" ]
