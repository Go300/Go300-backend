FROM golang:1.10.2

COPY ./src $GOPATH/go/src/Go300-backend/src
WORKDIR $GOPATH/go/src/Go300-backend/src

RUN go get github.com/smartystreets/goconvey/convey
RUN go get github.com/gorilla/mux
RUN go get github.com/mongodb/mongo-go-driver/mongo
