package main

import (
	_ "github.com/Oybek-uzb/posts_api_gateway/cmd/docs"
	"github.com/Oybek-uzb/posts_api_gateway/controller"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		posts := v1.Group("/posts")
		{
			posts.GET(":id", c.GetPost)
			posts.GET("", c.GetAllPosts)
			posts.PATCH(":id", c.UpdatePartialPost)
			posts.DELETE(":id", c.DeletePost)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")

	if err != nil {
		return
	}
}
