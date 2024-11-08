# Stage 1: Build the Go application
FROM golang:1.17 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o receipt-processor .

# Stage 2: Create a small image using only the binary
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/receipt-processor .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./receipt-processor"]

