package handlers

import (
	"adivii/forum-discussion/post/models"
	"adivii/forum-discussion/post/usecases"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type PostHttpHandler struct {
	postUsecase usecases.PostUsecase
}

func NewPostHttpHandler(postUsecase usecases.PostUsecase) PostHandler {
	return PostHttpHandler{
		postUsecase: postUsecase,
	}
}

// InsertPostData implements PostHandler.
func (p PostHttpHandler) InsertPostData(c echo.Context) error {
	reqBody := new(models.AddPostData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad Request")
	}

	if err := p.postUsecase.InsertForumData(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing Data Failed")
	}

	return response(c, http.StatusOK, "Success")
}
