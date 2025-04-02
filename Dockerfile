# Stage 1: Build
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rest-task-api

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/rest-task-api .
EXPOSE 8080
CMD ["./rest-task-api"]
