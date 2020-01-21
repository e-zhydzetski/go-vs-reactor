FROM golang:1.13.6-alpine3.11 AS builder
RUN apk update \
    && apk add --no-cache \
        git \
        ca-certificates \
        tzdata \
    && update-ca-certificates
RUN adduser -D -u 10001 appuser
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o build/go-service

# final stage
FROM alpine:3.11.3
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER appuser
WORKDIR /app
ENTRYPOINT ["./go-service"]
ENV PORT 8080
EXPOSE $PORT
COPY --from=builder /workspace/build/go-service ./go-service