# Use official Go image
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go modules and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN go build -o livebatch

# Run the binary
CMD ["./livebatch"]
