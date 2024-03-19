FROM golang:1.22 as builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux make build-http

FROM alpine:latest

WORKDIR /web
COPY --from=builder /tmp/bin/http ./
CMD ["./http"]
