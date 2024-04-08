# Use the official Go image as a base to build our application binary.
FROM golang:1.21.5 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies.
COPY go.mod ./
COPY go.sum ./

# Download the dependencies.
RUN go mod download

# Copy the entire project into the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o FinanxierApp ./cmd/app/http/main.go

# Start a new stage from an Alpine image to create a small final image
FROM alpine:latest  

# Add necessary packages.
RUN apk --no-cache add ca-certificates

# Set the working directory to root
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/FinanxierApp .

# Copy the entire config directory
COPY ./config /config

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./FinanxierApp"]
