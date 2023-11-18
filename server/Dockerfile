# Use the official golang image as base image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY . .

# Download and install any dependencies
RUN go mod download

RUN rm -f gocli.go

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o myapp



CMD ["./myapp"]







