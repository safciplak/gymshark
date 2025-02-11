FROM golang:1.23-alpine

WORKDIR /app

COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

EXPOSE 8081
CMD ["./main"] 