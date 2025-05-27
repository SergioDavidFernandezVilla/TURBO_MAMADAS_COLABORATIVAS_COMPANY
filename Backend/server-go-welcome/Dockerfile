# Use the official Go image as a base image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal image for the final stage
FROM alpine:latest

# Install ca-certificates for SSL support
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your application listens on
EXPOSE 8080

# Set the entrypoint to run the compiled binary
ENTRYPOINT ["./main"]