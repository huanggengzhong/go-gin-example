FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/huanggengzhong/go-gin-example
COPY . $GOPATH/src/github.com/huanggengzhong/go-gin-example
RUN go env
RUN go list -m all
RUN go build . 

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]