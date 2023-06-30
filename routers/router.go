package routers

import (
	"github.com/huanggengzhong/go-gin-example/pkg/setting"
	v1 "github.com/huanggengzhong/go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "成功",
		})
	})
	// 路由分组
	apiv1 := r.Group("/api/v1")
	{
		// 标签路由组
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// 文章路由组
		apiv1.GET("/articals", v1.GetArtical)
		apiv1.POST("/articals", v1.AddArtical)
		apiv1.PUT("/articals/:id", v1.EditArtical)
		apiv1.DELETE("/articals/:id", v1.DeleteArtical)
	}
	return r
}
