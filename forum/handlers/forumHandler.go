package handlers

import "github.com/labstack/echo"

type ForumHandler interface {
	InsertForumData(c echo.Context) error
}
