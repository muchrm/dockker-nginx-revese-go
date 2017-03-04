FROM golang:alpine
WORKDIR /go/src/app
RUN apk update \
    && apk upgrade \
    && apk add git
ENTRYPOINT go get -u github.com/gorilla/mux && go run main.go