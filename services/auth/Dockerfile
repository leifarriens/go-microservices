FROM golang:1.22.1-alpine AS build-stage

WORKDIR /app

COPY go.work ./
COPY go.work.sum ./

COPY ./internal ./internal
COPY ./services ./services

RUN go mod download

RUN go build -o /auth-service ./services/auth/cmd

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage ./auth-service ./auth-service

ENV ENVIRONMENT=container

EXPOSE 8080

CMD ["./auth-service"]
