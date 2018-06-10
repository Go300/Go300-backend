FROM golang:1.10.2

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin

COPY ./src /go/src/Go300-backend/src
WORKDIR /go/src/Go300-backend/src

RUN go get github.com/smartystreets/goconvey/convey
RUN go get github.com/gorilla/mux
RUN go get github.com/mongodb/mongo-go-driver/mongo
