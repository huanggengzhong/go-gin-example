package main

import (
	"fmt"
	"net/http"

	"github.com/huanggengzhong/go-gin-example/pkg/setting"
	"github.com/huanggengzhong/go-gin-example/routers"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
