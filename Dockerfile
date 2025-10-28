# Build stage
FROM golang:1.24-alpine AS builder
RUN apk --no-cache add build-base git mercurial gcc curl openssh-client musl-dev

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
ENV GOPATH=/go
ENV GOCACHE=/go/cache
ENV GOCACHE=/go-build
ENV GOMODCACHE=/go/pkg/mod
ENV CGO_ENABLED=0
ENV GOPRIVATE=github.com/PT-Sinarmas-Multifinance/
RUN git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"

WORKDIR /builder
COPY Makefile .
RUN make setup-doc
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make doc
RUN make build
RUN make setup-migration

# RUN State
FROM alpine:3.21
LABEL maintainer="SMMF Digital Team 2025"
WORKDIR /app
COPY --from=builder --chown=1001 /builder/bin/runner .
COPY --from=builder --chown=1001 /go/bin/goose .
RUN chmod +x /app/runner
COPY db/migrations ./db/migrations
COPY scripts ./scripts
COPY configs/config.yml.example ./configs/config.yml
RUN chmod +x scripts/startup/start.sh && chmod +x scripts/startup/wait-for.sh

RUN apk add --no-cache tzdata
ENV TZ="UTC"

# RUN /app/runner migrate
ENTRYPOINT ["/app/runner", "serve"]