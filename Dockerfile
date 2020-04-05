FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/kusole/go-gin-learn

COPY . $GOPATH/src/github.com/kusole/go-gin-learn

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./go-gin-learn"]
