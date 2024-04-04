# Start from the latest golang base image
FROM golang:1.21.4 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o service cmd/main.go
#RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o service cmd/common/main.go


# Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# This lib is required
RUN apk add gcompat
RUN apk add libc6-compat
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/service .

# Expose port 8080 to the outside world
EXPOSE 9000
# Test
# Command to run the executable
CMD ["./service"]