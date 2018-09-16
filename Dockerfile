FROM golang:alpine

# Copy sources
COPY . /go/src/github.com/ravenq/gvf-server
WORKDIR /go/src/github.com/ravenq/gvf-server

# build
RUN go build

# launches gvf-server
CMD ["./gvf-server"]