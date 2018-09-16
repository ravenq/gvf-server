FROM golang:alpine

# Copy sources
COPY . /go/src/gitlab.com/ravenq/gvf-server
WORKDIR /go/src/gitlab.com/ravenq/gvf-server

# build
RUN go build

# launches gvf-server
CMD ["./gvf-server"]