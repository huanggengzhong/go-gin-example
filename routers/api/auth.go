package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/go-gin-example/models"
	"github.com/huanggengzhong/go-gin-example/pkg/e"
	"github.com/huanggengzhong/go-gin-example/pkg/logging"
	"github.com/huanggengzhong/go-gin-example/pkg/util"
)

type auth struct {
	Username string `valid:Required;MaxSize(50)`
	Password string `valid:Required;MaxSize(50)`
}

// ShowAccount godoc
//
//	@Summary		登录
//	@Description	获取token
//	@Tags			用户
//	@Accept			json
//	@Produce		json
//	@Param			username	path		string	true	"用户名"
//	@Param			password	path		string	true	"密码"
//	@Success		200 {string} json "{"code":200,"data":{"token":"xx"},"msg":"ok"}"
//	@Router			/api/v1/auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	valid := validation.Validation{}

	// 校验
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")

	a := auth{Username: username, Password: password}

	ok, _ := valid.Valid(&a)

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			// log.Printf("err.Key:%s,err.Message:%s", err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
