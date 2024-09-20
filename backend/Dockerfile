# Stage 1: Build the Go binary with CGO enabled
FROM golang:1.23.1-alpine AS build

# Install build dependencies for SQLite
RUN apk add --no-cache gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Set CGO_ENABLED to 1 for SQLite
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go app with CGO enabled
RUN go build -o backend-vid-dl ./cmd/vid-dl

# Stage 2: Run the Go binary in a minimal image
FROM alpine:latest

# Install SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=build /app/backend-vid-dl .

# Copy the bin folder (for yt-dlp)
COPY --from=build /app/bin ./bin

# Expose port 3000
EXPOSE 3000

# Command to run the binary
CMD ["./backend-vid-dl"]
