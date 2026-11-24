FROM golang:1.26.5-alpine
# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy go mod and sum files
COPY . .

# Expose the port
EXPOSE 8080

RUN go build -o main ./cmd/main.go

# Run the application
CMD ["./main"]