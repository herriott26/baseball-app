# Start from golang v1.11 base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Joe Herriott <herriott26@gmail.com>"
LABEL app="Transaction Aggregator"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/kepler26/baseball-app/aggregator

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/aggregator .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/aggregator .

CMD ["./aggregator"] 