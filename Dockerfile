FROM golang:1.7
RUN mkdir -p /test
RUN go get github.com/julienschmidt/httprouter
WORKDIR /test
ADD . /test
RUN go build ./test.go
CMD ["./test"]
