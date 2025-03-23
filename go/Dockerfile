FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/server cmd/sso/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /app/server /app/server
COPY --from=builder /app/storage /app/storage
COPY --from=builder /app/config /app/config

EXPOSE 44044

CMD ["/app/server", "--config=./app/config/dev.yml"]