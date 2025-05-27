# Use the official Go image as a base image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a minimal base image for the final image
FROM alpine:latest

# Install ca-certificates to ensure HTTPS works
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the application listens on
EXPOSE 8082

# Set the entrypoint to run the compiled binary
ENTRYPOINT ["./main"]