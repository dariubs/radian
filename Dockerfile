FROM golang:1.9

ADD . .

# Install dependencies
RUN go get -v -d ./...

# Build the app command inside the container.
RUN make

# Run the app command by default when the container starts.
ENTRYPOINT /go/app

# Document that the service listens on port 2112.
EXPOSE 2112
