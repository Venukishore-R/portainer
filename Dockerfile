# Use official Golang image as base
FROM golang:1.22-alpine

# Set working directory inside container
WORKDIR /app

# Copy files
COPY . .

# Build the Go application
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
