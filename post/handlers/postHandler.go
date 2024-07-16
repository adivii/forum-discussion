package handlers

import (
	"github.com/labstack/echo"
)

type PostHandler interface {
	InsertPostData(c echo.Context) error
	GetAllOriginalPost(c echo.Context) error
	GetPostById(c echo.Context) error
}
