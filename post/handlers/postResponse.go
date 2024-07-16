package handlers

import (
	"adivii/forum-discussion/post/entities"

	"github.com/labstack/echo"
)

type BaseResponse struct {
	Message string `json:"message"`
}

type MultiPostResponse struct {
	Message string                     `json:"message"`
	Data    []entities.PostResponseDto `json:"data"`
}

func response(c echo.Context, responseCode int, message string) error {
	return c.JSON(responseCode, &BaseResponse{
		Message: message,
	})
}

func responseMultiplePost(c echo.Context, responseCode int, message string, data []entities.PostResponseDto) error {
	return c.JSON(responseCode, &MultiPostResponse{
		Message: message,
		Data:    data,
	})
}
