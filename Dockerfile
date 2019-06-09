FROM golang

ADD . /go/src/github.com/leepuppychow/employees_API
WORKDIR /go/src/github.com/leepuppychow/employees_API

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers
RUN go get github.com/go-sql-driver/mysql

RUN go install github.com/leepuppychow/employees_API

ENV GO_TEST_USER=root
ENV GO_TEST_PASS=password
ENV GO_TEST_HOST=localhost
ENV GO_TEST_PORT=8080

ENTRYPOINT /go/bin/employees_API

EXPOSE 8080