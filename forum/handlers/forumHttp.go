package handlers

import (
	"adivii/forum-discussion/forum/models"
	"adivii/forum-discussion/forum/usecases"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ForumHttpHandler struct {
	forumUsecase usecases.ForumUsecase
}

func NewForumHttpHandler(forumUsecase usecases.ForumUsecase) ForumHandler {
	return ForumHttpHandler{
		forumUsecase: forumUsecase,
	}
}

// InsertForumData implements ForumHandler.
func (f ForumHttpHandler) InsertForumData(c echo.Context) error {
	reqBody := new(models.AddForumData)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad Request")
	}

	if err := f.forumUsecase.InsertForumData(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing Data Failed")
	}

	return response(c, http.StatusOK, "Success")
}
