FROM golang:1.18

WORKDIR /go/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /go/radian

# Run the app command by default when the container starts.
ENTRYPOINT /go/radian

# Document that the service listens on port 2112.
EXPOSE 2112
