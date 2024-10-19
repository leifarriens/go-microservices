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

RUN go install github.com/cosmtrek/air@v1.43.0

COPY ./services/ping ./services/ping

WORKDIR /app/services/ping

CMD ["air", "-c", ".air.toml"]

FROM build-base AS build-prod

COPY ./services/ping ./services/ping

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s' -x -o /ping-service ./services/ping/cmd

FROM gcr.io/distroless/base-debian11 AS prod

WORKDIR /

COPY --from=build-prod ./ping-service ./ping-service

ENV ENVIRONMENT=container

EXPOSE 1324

CMD ["./ping-service"]