FROM golang:1.22.1-alpine AS build-base

WORKDIR /app 

COPY go.work go.work.sum ./
COPY ./internal/shared/go.mod ./internal/shared/go.sum ./internal/shared/
COPY ./services/auth/go.mod ./services/auth/go.sum ./services/auth/
COPY ./services/product/go.mod ./services/product/go.sum ./services/product/

RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

FROM build-base AS dev

RUN go install github.com/cosmtrek/air@v1.43.0 && \
  go install github.com/swaggo/swag/cmd/swag@v1.16.4

COPY ./internal ./internal
COPY ./services ./services

WORKDIR /app/services/product

CMD ["air", "-c", ".air.toml"]

FROM build-base AS build-prod

COPY ./internal ./internal
COPY ./services ./services

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s' -x -o /product-service ./services/product/cmd

FROM gcr.io/distroless/base-debian11 AS prod

WORKDIR /

COPY --from=build-prod ./product-service ./product-service

ENV ENVIRONMENT=container

EXPOSE 1324

CMD ["./product-service"]