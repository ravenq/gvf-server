FROM golang:1.11-alpine

MAINTAINER ravenq <flw_dream@126.com>

# Copy sources
COPY . /go/src/github.com/ravenq/gvf-server
WORKDIR /go/src/github.com/ravenq/gvf-server

EXPOSE 8080

# build
RUN go build

# launches gvf-server
CMD ["./gvf-server"]
