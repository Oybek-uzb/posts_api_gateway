package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	pbp "github.com/Oybek-uzb/posts_api_gateway/pkg/api/posts_crud_service"
	"github.com/gin-gonic/gin"
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
func (c *Controller) GetPost(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	postId, err2 := strconv.Atoi(gc.Param("id"))

	if err2 != nil {
		gc.JSON(http.StatusInternalServerError, err2)
		return
	}

	response, err := c.Services.PostsCRUDService().GetPost(ctx, &pbp.GetPostRequest{
		Id: int32(postId),
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
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
func (c *Controller) GetAllPosts(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := c.Services.PostsCRUDService().GetAllPosts(ctx, &pbp.GetAllPostsRequest{})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
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
func (c *Controller) UpdatePartialPost(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	postId, err := strconv.Atoi(gc.Param("id"))

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	var data pbp.Post

	err = gc.ShouldBindJSON(&data)

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	data.Id = int32(postId)

	response, err := c.Services.PostsCRUDService().UpdatePartialPost(ctx, &pbp.UpdatePartialPostRequest{
		UpdateData: &data,
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
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
func (c *Controller) DeletePost(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	postId, err := strconv.Atoi(gc.Param("id"))

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := c.Services.PostsCRUDService().DeletePost(ctx, &pbp.DeletePostRequest{
		Id: int32(postId),
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}
