FROM golang:1.10.2

COPY ./src /go/src/Go300-backend/src
WORKDIR /go/src/Go300-backend/src
RUN export GOROOT=/usr/local/go
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:$GOROOT/bin

RUN cat Makefile
RUN make && echo "ok"

RUN go get github.com/smartystreets/goconvey/convey
RUN go get github.com/gorilla/mux
RUN go get github.com/mongodb/mongo-go-driver/mongo
