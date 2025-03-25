# Use the official Golang image to build the application
FROM golang:1.21 as builder

# Set the working directory
WORKDIR /app

# Copy the Go source code
COPY p3.go .

# Build the Go application
RUN go mod init datetime-app && go mod tidy && go build -o app

# Use a smaller base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled application from the builder stage
COPY --from=builder /app/app .

# Expose the required port
EXPOSE 8080

# Run the application
CMD ["./app"]
