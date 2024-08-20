# Use the official Golang image as a base image
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Install Prisma client
RUN go install github.com/steebchen/prisma-client-go@latest

# Generate Prisma client
RUN go run github.com/steebchen/prisma-client-go generate

# Build the Go application
RUN go build -o main .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]
