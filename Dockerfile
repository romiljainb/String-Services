FROM golang:1.7
RUN mkdir -p /testString
RUN go get github.com/gorilla/mux
WORKDIR /testString
ADD . /testString
RUN go build ./testString.go
CMD ["./testString"]
