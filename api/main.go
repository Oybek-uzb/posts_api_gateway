package api

import (
	"github.com/Oybek-uzb/posts_api_gateway/services"
	"github.com/gin-gonic/gin"

	_ "github.com/Oybek-uzb/posts_api_gateway/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Iman test-task API
// @version         1.0
// @description     This is a sample API Gateway for CRUD of posts.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func NewServer(sm services.ServiceManager) *gin.Engine {

	r := gin.Default()

	c := NewController(sm)

	posts := r.Group("/api/v1")
	{
		posts.POST("/posts/fetch-from-remote", c.FetchFromRemote)
		posts.GET("/posts", c.GetAllPosts)
		posts.GET("/posts/:id", c.GetPost)
		posts.PATCH("/posts/:id", c.UpdatePartialPost)
		posts.DELETE("/posts/:id", c.DeletePost)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
