FROM golang:1.16-alpine

WORKDIR $GOPATH/src/go-prac-site
COPY . $GOPATH/src/go-prac-site

RUN go build -o ./app

EXPOSE 8080

CMD ["./app"]
