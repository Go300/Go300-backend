FROM golang:1.10.2

COPY ./src /go/src/Go300-backend/src
WORKDIR /go/src/Go300-backend/src

RUN go get github.com/smartystreets/goconvey/convey
#RUN go install github.com/smartystreets/goconvey
#CMD /go/bin/goconvey -host 0.0.0.0

RUN go get github.com/gorilla/mux
RUN go get github.com/go-bongo/bongo
RUN go get github.com/jasonlvhit/gocron
RUN go get github.com/subosito/gotenv
