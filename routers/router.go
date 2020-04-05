package routers

import (
	"github.com/gin-gonic/gin"
	setting "github.com/kusole/go-gin-learn/pkg"
	"github.com/kusole/go-gin-learn/routers/api"
	v1 "github.com/kusole/go-gin-learn/routers/api/v1"
	jwt "github.com/kusole/go-gin-learn/middleware"
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

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		// 新建标签
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

// tag
/*
	GET: http://127.0.0.1:8000/api/v1/tags?name=2&state=1&page=1
	POST: http://127.0.0.1:8000/api/v1/tags?name=2&state=1&created_by=test
	PUT: http://127.0.0.1:8000/api/v1/tags/1?name=linux&modified_by=liutao
	DELETE: http://127.0.0.1:8000/api/v1/tags/1
*/

// article
/*
	POST：http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1
	GET：http://127.0.0.1:8000/api/v1/articles
	GET：http://127.0.0.1:8000/api/v1/article/1
	PUT：http://127.0.0.1:8000/api/v1/article/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0
	DELETE：http://127.0.0.1:8000/api/v1/article/1
*/
