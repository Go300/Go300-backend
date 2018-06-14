FROM iron/go:dev

WORKDIR /src
ENV SRC_DIR=/go/src/github.com/Go300/Go300-backend/
ADD ./src $SRC_DIR
RUN cd $SRC_DIR; go build -o myapp; cp myapp /src/

RUN go get github.com/smartystreets/goconvey/convey
RUN go get github.com/gorilla/mux
RUN go get github.com/mongodb/mongo-go-driver/mongo
