FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/osuzhdayu-bot .

FROM alpine:3.21

RUN apk add --no-cache ca-certificates
WORKDIR /app
ENV CONFIG_PATH=/app/config.cfg

# config.cfg is intentionally not copied into the image. Mount it at runtime.
COPY --from=builder /out/osuzhdayu-bot ./osuzhdayu-bot

CMD ["/app/osuzhdayu-bot"]
