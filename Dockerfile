FROM golang:alpine as builder

ENV GOBIN /go/bin
ENV GOPATH /app
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p /app/src/krdv-test
WORKDIR /app/src/krdv-test

ADD . /app/src/krdv-test

RUN apk update && \
apk add git && \
rm -rf /tmp/* && \
rm -rf /var/cache/apk/* && \
go get .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -i -o main main.go
EXPOSE 8080
CMD ["./main"]

