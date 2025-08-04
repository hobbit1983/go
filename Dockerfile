############################
# STEP 1 build executable binary
############################
FROM golang AS builder
WORKDIR $GOPATH/src/github.com/hobbit1983/go
COPY . /go/src/github.com/hobbit1983/go

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=auto
ENV GOPATH=/go

WORKDIR $GOPATH/src/github.com/hobbit1983/go
COPY . /go/src/github.com/hobbit1983/go

RUN go build -o /go/bin/run /go/src/github.com/hobbit1983/go/main.go

############################
# STEP 2 build a small image
############################
FROM alpine
COPY --from=builder /go/bin/run /go/bin/run
ENTRYPOINT ["/go/bin/run"]
