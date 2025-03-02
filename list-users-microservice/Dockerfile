# Use Golang official image
FROM golang:1.19-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create a directory for the app
WORKDIR /app

# Copy the Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 2003

# Command to run the executable
CMD ["./main"]
