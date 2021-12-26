package httpv1

import (
	"golang-template-module/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserGet godoc
// @Summary      Get user
// @Description  Get user by id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Success      200  {object}  models.UserBaseModel
// @Router       /user/{id} [get]
func RouterUserGetId(c echo.Context) error {
	id := c.Param("id")
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Only uint")
	}
	user, err := service.GetUserService(u64)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// UserGetList godoc
// @Summary      Get user list
// @Description  Get user list
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserListBaseModel
// @Router       /user/list [get]
func RouterUserGetList(c echo.Context) error {
	users, err := service.GetUserListService()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
