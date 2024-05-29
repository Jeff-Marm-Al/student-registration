# Use a Go runtime image as the base image
FROM golang:latest as builder

# Set the working directory in the builder stage
RUN mkdir /app
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download and install Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o student-registration-api

# Use a lightweight Alpine image as the base image for the final container
FROM ubuntu:latest

# Copy the built executable from the builder stage into the final container
COPY --from=builder /app/student-registration-api /app/
WORKDIR /app

# Set execute permissions for the executable
RUN chmod +x /app/student-registration-api

# Expose the port on which your API server runs
EXPOSE 8080

# Command to run the executable
CMD ["./student-registration-api"]
