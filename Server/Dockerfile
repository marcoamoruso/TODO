FROM golang:latest

ENV GOPATH /go
RUN mkdir -p "$GOPATH/src/TODO"

ADD . ${GOPATH}/src/TODO/
WORKDIR ${GOPATH}/src/TODO/

RUN go mod init TODO
RUN go mod tidy