FROM golang:alpine

# Copy sources
COPY . /go/src/gitlab.com/ravenq/gvf-server
WORKDIR /go/src/gitlab.com/ravenq/gvf-server

# launches gvf-server
CMD ["./gvf-server"]