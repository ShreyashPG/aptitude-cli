# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o aptitude-app .

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/aptitude-app .
COPY Aptitude.json .  

ENTRYPOINT ["./aptitude-app"]
