package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"echo-swaggo-example/pkg/httputil"
	"echo-swaggo-example/pkg/model"
	"github.com/labstack/echo/v4"
)

// ShowAccount godoc
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx echo.Context)  error {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	account, err := model.AccountOne(aid)
	if err != nil {
		return httputil.NewError(ctx, http.StatusNotFound, err)

	}
	return ctx.JSON(http.StatusOK, account)
}

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q" Format(email)
// @Success 200 {array} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [get]
func (c *Controller) ListAccounts(ctx echo.Context)  error {
	q := ctx.QueryParam("q")
	accounts, err := model.AccountsAll(q)
	if err != nil {
		return  httputil.NewError(ctx, http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, accounts)
}

// AddAccount godoc
// @Summary Add an account
// @Description add by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body model.AddAccount true "Add account"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [post]
func (c *Controller) AddAccount(ctx echo.Context)  error {
	var addAccount model.AddAccount
	if err := ctx.Bind(&addAccount); err != nil {
		return  httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	if err := addAccount.Validation(); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	account := model.Account{
		Name: addAccount.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	account.ID = lastID
	return ctx.JSON(http.StatusOK, account)
}

// UpdateAccount godoc
// @Summary Update an account
// @Description Update by json account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID"
// @Param  account body model.UpdateAccount true "Update account"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [patch]
func (c *Controller) UpdateAccount(ctx echo.Context)  error {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	var updateAccount model.UpdateAccount
	if err := ctx.Bind(&updateAccount); err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	account := model.Account{
		ID:   aid,
		Name: updateAccount.Name,
	}
	err = account.Update()
	if err != nil {
		return httputil.NewError(ctx, http.StatusNotFound, err)

	}
	return ctx.JSON(http.StatusOK, account)
}

// DeleteAccount godoc
// @Summary Delete an account
// @Description Delete by account ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param  id path int true "Account ID" Format(int64)
// @Success 204 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [delete]
func (c *Controller) DeleteAccount(ctx echo.Context)  error {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	err = model.Delete(aid)
	if err != nil {
		return httputil.NewError(ctx, http.StatusNotFound, err)

	}
	return ctx.JSON(http.StatusNoContent,`{}`)
}

// UploadAccountImage godoc
// @Summary Upload account image
// @Description Upload file
// @Tags accounts
// @Accept  multipart/form-data
// @Produce  json
// @Param  id path int true "Account ID"
// @Param file formData file true "account image"
// @Success 200 {object} controller.Message
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id}/images [post]
func (c *Controller) UploadAccountImage(ctx echo.Context)  error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	file, err := ctx.FormFile("file")
	if err != nil {
		return httputil.NewError(ctx, http.StatusBadRequest, err)

	}
	return ctx.JSON(http.StatusOK, Message{Message: fmt.Sprintf("upload complete userID=%d filename=%s", id, file.Filename)})
}
