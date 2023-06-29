package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/go-gin-example/models"
	"github.com/huanggengzhong/go-gin-example/pkg/e"
	"github.com/huanggengzhong/go-gin-example/pkg/setting"
	"github.com/huanggengzhong/go-gin-example/pkg/util"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data := make(map[string]interface{})

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func AddTag(c *gin.Context) {

}
func EditTag(c *gin.Context) {

}
func DeleteTag(c *gin.Context) {

}
