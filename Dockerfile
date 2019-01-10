FROM golang:alpine
ADD . /go/src/github.com/c0z0/go-shrt
RUN go install github.com/c0z0/go-shrt
CMD ["/go/bin/go-shrt"]
EXPOSE 3000
