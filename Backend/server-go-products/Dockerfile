# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application listens on
EXPOSE 8081

# Command to run the executable
CMD ["./main"]