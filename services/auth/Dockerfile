FROM golang:1.22.1-alpine AS build-base

WORKDIR /app 

COPY go.work go.work.sum ./
COPY ./internal/shared/go.mod ./internal/shared/go.sum ./internal/shared/
COPY ./services/auth/go.mod ./services/auth/go.sum ./services/auth/
COPY ./services/product/go.mod ./services/product/go.sum ./services/product/
COPY ./services/ping/go.mod ./services/ping/

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

FROM build-base AS dev

RUN go install github.com/cosmtrek/air@v1.43.0 && \
  go install github.com/swaggo/swag/cmd/swag@v1.16.4

COPY ./internal ./internal
COPY ./services ./services

WORKDIR /app/services/auth

CMD ["air", "-c", ".air.toml"]

FROM build-base AS build-prod

COPY ./internal ./internal
COPY ./services ./services

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s' -x -o /auth-service ./services/auth/cmd

FROM gcr.io/distroless/base-debian11 AS prod

WORKDIR /

COPY --from=build-prod ./auth-service ./auth-service

ENV ENVIRONMENT=container

EXPOSE 1323

CMD ["./auth-service"]