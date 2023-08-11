package routers

import (
	"net/http"

	"github.com/huanggengzhong/go-gin-example/pkg/upload"

	"github.com/gin-gonic/gin"

	docs "github.com/huanggengzhong/go-gin-example/docs"

	"github.com/huanggengzhong/go-gin-example/middleware/jwt"
	"github.com/huanggengzhong/go-gin-example/pkg/setting"
	"github.com/huanggengzhong/go-gin-example/routers/api"
	v1 "github.com/huanggengzhong/go-gin-example/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	// // programatically set swagger info
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	// 静态文件访问
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//获取token
	r.GET("/auth", api.GetAuth)
	// swag
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 上传
	r.POST("/upload", api.UploadImage)

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
