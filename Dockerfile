FROM alpine

RUN apk add --no-cache ca-certificates

# Copy the binary file and set it as entrypoint
ADD go-elastic-content-service /
ENTRYPOINT ["/go-elastic-content-service"]

# The service listens on port 8080 by default.
EXPOSE 8080