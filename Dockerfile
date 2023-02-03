FROM golang:1.19.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8001

CMD ["/app/server"]