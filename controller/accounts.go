package controller

import (
	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  models.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
}

// ListAccounts godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   models.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
}

// AddAccount godoc
// @Summary      Add an account
// @Description  add by json account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account  body      models.AddAccount  true  "Add account"
// @Success      200      {object}  models.Account
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
}

// UpdateAccount godoc
// @Summary      Update an account
// @Description  Update by json account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "Account ID"
// @Param        account  body      models.UpdateAccount  true  "Update account"
// @Success      200      {object}  models.Account
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /accounts/{id} [patch]
func (c *Controller) UpdateAccount(ctx *gin.Context) {
}

// DeleteAccount godoc
// @Summary      Delete an account
// @Description  Delete by account ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"  Format(int64)
// @Success      204  {object}  models.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [delete]
func (c *Controller) DeleteAccount(ctx *gin.Context) {
}

// UploadAccountImage godoc
// @Summary      Upload account image
// @Description  Upload file
// @Tags         accounts
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path      int   true  "Account ID"
// @Param        file  formData  file  true  "account image"
// @Success      200   {object}  controller.Message
// @Failure      400   {object}  httputil.HTTPError
// @Failure      404   {object}  httputil.HTTPError
// @Failure      500   {object}  httputil.HTTPError
// @Router       /accounts/{id}/images [post]
func (c *Controller) UploadAccountImage(ctx *gin.Context) {
}
