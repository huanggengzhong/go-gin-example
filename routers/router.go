package routers

import (
	"github.com/huanggengzhong/go-gin-example/middleware/jwt"
	"github.com/huanggengzhong/go-gin-example/pkg/setting"
	"github.com/huanggengzhong/go-gin-example/routers/api"
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

	//获取token
	r.GET("/auth", api.GetAuth)
	// 路由分组
	apiv1 := r.Group("/api/v1")

	//token验证
	apiv1.Use(jwt.JWT())
	{
		// 标签路由组
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// 文章路由组
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
