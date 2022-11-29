FROM golang:alpine AS builder
WORKDIR /build

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

COPY . .
RUN go build std
RUN go build -trimpath -ldflags '-w -s -extldflags "-static"' -o /build/app cmd/hello-world-server/main.go

FROM ghcr.io/ironpeakservices/iron-scratch/iron-scratch:1.0.0
COPY --from=builder /build/app /app

EXPOSE 8080
ENTRYPOINT ["/app"]