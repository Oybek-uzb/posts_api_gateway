package controller

import (
	"github.com/Oybek-uzb/posts_api_gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPost godoc
// @Summary      Get a post
// @Description  Get by post ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  models.Post
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts/{id} [get]
func (c *Controller) GetPost(ctx *gin.Context) {
	post := models.Post{
		ID:     11,
		UserId: 11,
		Title:  "Liboy",
		Body:   "Liboy",
	}
	ctx.JSON(http.StatusOK, post)
}

// GetAllPosts godoc
// @Summary      List posts
// @Description  Get all posts
// @Tags         posts
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Post
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts [get]
func (c *Controller) GetAllPosts(ctx *gin.Context) {
}

// UpdatePartialPost godoc
// @Summary      Update a post
// @Description  Partial update a post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "Post ID"
// @Param        post  body      models.UpdatePost  true  "Update post"
// @Success      200      {object}  models.Post
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /posts/{id} [patch]
func (c *Controller) UpdatePartialPost(ctx *gin.Context) {
}

// DeletePost godoc
// @Summary      Delete a post
// @Description  Delete by post ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      204  {object}  models.Post
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /posts/{id} [delete]
func (c *Controller) DeletePost(ctx *gin.Context) {
}
