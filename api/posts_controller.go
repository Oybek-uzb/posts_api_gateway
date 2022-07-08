package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	pbpc "github.com/Oybek-uzb/posts_api_gateway/pkg/api/posts_crud_service"
	pbp "github.com/Oybek-uzb/posts_api_gateway/pkg/api/posts_service"
)

// FetchFromRemote godoc
// @Summary      Fetch posts from remote origin
// @Description  Fetch just by clicking
// @Tags         posts
// @Accept       json
// @Produce      json
// @Success      200  {object}  bool
// @Failure      400      {object}  HTTPError
// @Failure      404      {object}  HTTPError
// @Failure      500      {object}  HTTPError
// @Router       /posts/fetch-from-remote [post]
func (c *Controller) FetchFromRemote(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := c.Services.RemotePostsService().GetRemotePosts(ctx, &pbp.GetRemotePostsRequest{})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}

// GetPost godoc
// @Summary      Get a post
// @Description  Get by post ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  Post
// @Failure      400      {object}  HTTPError
// @Failure      404      {object}  HTTPError
// @Failure      500      {object}  HTTPError
// @Router       /posts/{id} [get]
func (c *Controller) GetPost(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	postId, err2 := strconv.Atoi(gc.Param("id"))

	if err2 != nil {
		gc.JSON(http.StatusInternalServerError, err2)
		return
	}

	response, err := c.Services.PostsCRUDService().GetPost(ctx, &pbpc.GetPostRequest{
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
// @Accept       */*
// @Produce      json
// @Success      200  {array}   Post
// @Failure      400      {object}  HTTPError
// @Failure      404      {object}  HTTPError
// @Failure      500      {object}  HTTPError
// @Router       /posts [get]
func (c *Controller) GetAllPosts(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	response, err := c.Services.PostsCRUDService().GetAllPosts(ctx, &pbpc.GetAllPostsRequest{})

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
// @Param        post  body      UpdatePost  true  "Update post"
// @Success      200      {object}  Post
// @Failure      400      {object}  HTTPError
// @Failure      404      {object}  HTTPError
// @Failure      500      {object}  HTTPError
// @Router       /posts/{id} [patch]
func (c *Controller) UpdatePartialPost(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	postId, err := strconv.Atoi(gc.Param("id"))

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	var data pbpc.Post

	err = gc.ShouldBindJSON(&data)

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	data.Id = int32(postId)

	response, err := c.Services.PostsCRUDService().UpdatePartialPost(ctx, &pbpc.UpdatePartialPostRequest{
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
// @Success      204  {object}  Post
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /posts/{id} [delete]
func (c *Controller) DeletePost(gc *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(7))
	defer cancel()

	postId, err := strconv.Atoi(gc.Param("id"))

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := c.Services.PostsCRUDService().DeletePost(ctx, &pbpc.DeletePostRequest{
		Id: int32(postId),
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, err)
		return
	}

	c.ResponseProtoJson(gc, response)
}

type Post struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type UpdatePost struct {
	Title  string `json:"title"`
	UserId int    `json:"user_id"`
	Body   string `json:"body"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
