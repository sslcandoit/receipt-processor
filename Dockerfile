# Use the official Golang image as the builder
FROM golang:1.17 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and go sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o receipt-processor

# Use a smaller base image for the final container
FROM alpine:latest

# Set the working directory in the final container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/receipt-processor .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./receipt-processor"]
