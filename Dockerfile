# Use official lightweight Go image
FROM golang:1.24-alpine AS builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev sqlite-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/server

FROM alpine:latest as runner
ENV GO_ENV=production
WORKDIR /app
RUN apk add --no-cache sqlite-libs
COPY --from=builder /app/main /main
EXPOSE 8080
CMD ["/main"]
