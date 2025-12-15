FROM golang:1.25.5-alpine AS builder

ENV GOTOOLCHAIN=local
WORKDIR /app
    
RUN apk add --no-cache git
    
COPY go.mod go.sum ./
RUN go mod download
    
COPY . .
    
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
        -o server ./cmd/server
    
FROM alpine:latest
    
WORKDIR /app
    
COPY --from=builder /app/server .
    
EXPOSE 3000
    
CMD ["./server"]
    