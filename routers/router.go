package routers

import (
	"github.com/gin-gonic/gin"
	setting "github.com/kusole/go-gin-learn/pkg"
	v1 "github.com/kusole/go-gin-learn/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiv1 := r.Group("api/v1")
	{
		// 获取标签列表
		// GET http://127.0.0.1:8000/api/v1/tags?name=2&state=1&page=1
		apiv1.GET("/tags", v1.GetTags)

		// 新建标签
		// POST http://127.0.0.1:8000/api/v1/tags?name=2&state=1&created_by=test
		apiv1.POST("/tags", v1.AddTag)

		// 指定更新标签
		apiv1.PUT("/tags/:id", v1.EditTag)

		// 指定删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)

		// 获取文章
		apiv1.GET("/article/:id", v1.GetArticle)

		// 新建文章
		apiv1.POST("/articles", v1.AddArticle)

		// 更新指定文章
		apiv1.PUT("/article/:id", v1.EditArticle)

		// 删除指定文章
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
