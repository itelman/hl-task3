# Use the official Golang image to build the Go app
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app, assuming the binary will be named 'main'
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /

# Copy the built binary from the previous stage
COPY --from=builder /main .
COPY --from=builder . .

# Command to run the executable
CMD ["./main"]
