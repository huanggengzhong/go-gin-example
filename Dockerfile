FROM scratch

WORKDIR $GOPATH/src/github.com/huanggengzhong/go-gin-example
COPY . $GOPATH/src/github.com/huanggengzhong/go-gin-example

EXPOSE 8000
CMD ["./go-gin-example"]