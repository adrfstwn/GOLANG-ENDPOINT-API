# Use golang:alpine as the base image
FROM golang:alpine

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

# Create a temporary directory for downloading dependencies
RUN mkdir /tmp/gocache

# Make the /app directory temporarily read-write
RUN chmod -R +w /app

# Set the GOPATH to /tmp/gocache
ENV GOPATH /tmp/gocache

# Install build dependencies
RUN apk add --no-cache git

# Download and install dependencies using go mod
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o uas-api

# Make the /app directory read-only again
RUN chmod -R -w /app

# Expose the port used by the application
EXPOSE 1234

# Command to run the application when the container starts
CMD ["./uas-api"]
