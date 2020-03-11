# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Inject IT <contact@injectit.io>"
ENV GOPRIVATE="github.com/lukaszozimek"
RUN git config --global url."https://eea843361a185d3b59a24b03d42c3724ba64eb15:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8088
# Grpc server
EXPOSE 8086

# Command to run the executable
CMD ["./main"]