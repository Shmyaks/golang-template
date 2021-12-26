package models

import "github.com/labstack/echo/v4"

var (
	NotFoundError        = echo.NewHTTPError(404, "Not Found")
	ConnectionError      = echo.NewHTTPError(503, "Connection Error")
	ExistFoundError      = echo.NewHTTPError(409, "Item is already exists")
	PasswordOrLoginError = echo.NewHTTPError(403, "Password or Login is incorrect")
)
