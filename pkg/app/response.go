package app

import (
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Responser(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"data": data,
		"msg":  e.GetMsg(errCode),
	})
}
