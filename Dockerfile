FROM golang:1.8.3

RUN mkdir -p /go/src/github.com/furkhat/k8s-users

WORKDIR /go/src/github.com/furkhat/k8s-users

ADD . /go/src/github.com/furkhat/k8s-users

RUN go get ./webapp

RUN go build -i -o ./app ./webapp/main.go

EXPOSE 8080

CMD ["./app"]
