FROM golang:1.22-alpine AS dev

WORKDIR /app
RUN apk add --no-cache git && go install github.com/air-verse/air@v1.52.3
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]

FROM golang:1.22-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /chat-server ./cmd/chat-server

FROM alpine:3.20 AS runtime

WORKDIR /app
COPY --from=build /chat-server /app/chat-server
EXPOSE 8080
CMD ["/app/chat-server"]
