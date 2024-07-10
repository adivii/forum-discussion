package handlers

import "github.com/labstack/echo"

type PostHandler interface {
	InsertPostData(c echo.Context) error
}
