FROM golang:1.22.4-bookworm AS build-stage

WORKDIR /app

# Установка зависимостей для CGO
RUN apt-get update && apt-get install -y gcc

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o main ./cmd

FROM debian:bookworm-slim AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/main /card-detect-demo
COPY ./config ./config
COPY ./lib ./lib
COPY ./models ./models
COPY ./template ./template

RUN mkdir -p ./storage
RUN mkdir -p ./tmp

EXPOSE 8080

RUN chmod +x /card-detect-demo

ENTRYPOINT ["/card-detect-demo"]
