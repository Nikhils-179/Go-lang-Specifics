FROM golang:1.21.3-alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /dependency-inversion-example

# Use a smaller base image for running the application
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the pre-built binary from the build stage
COPY --from=build /dependency-inversion-example .

# Command to run the application
CMD ["./dependency-inversion-example"]
