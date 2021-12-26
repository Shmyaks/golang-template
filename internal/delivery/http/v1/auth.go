package httpv1

import (
	"fmt"
	"golang-template-module/internal/delivery/handlers"
	"golang-template-module/internal/models"
	"golang-template-module/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Auth godoc
// @Summary      Auth user
// @Description  auth User
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Auth   body   models.UserAuthModel  true "Auntification"
// @Success      200  {object}  models.UserTokenModel
// @Router       /auth [post]
func RouterAuth(c echo.Context) (err error) {
	user := new(models.UserAuthModel)
	if err = handlers.ValidateModel(c, user); err != nil {
		return err
	}
	tokenModel, err := service.Authorization(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tokenModel)
}

// CreateUser godoc
// @Summary      Create user
// @Description  create user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        Create   body   models.UserAuthModel  true "Creating"
// @Success      200  {object}  models.UserTokenModel
// @Router       /create [post]
func RouterRegistration(c echo.Context) (err error) {
	user := new(models.UserAuthModel)
	if err = handlers.ValidateModel(c, user); err != nil {
		fmt.Print(user)
		return err
	}
	tokenModel, err := service.CreateUser(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tokenModel)
}
