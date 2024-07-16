package handlers

import (
	"adivii/forum-discussion/post/models"
	"adivii/forum-discussion/post/usecases"
	"net/http"
	"strconv"

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

// GetAllOriginalPost implements PostHandler.
func (p PostHttpHandler) GetAllOriginalPost(c echo.Context) error {
	data, err := p.postUsecase.GetAllOriginalPost()
	if err != nil {
		return response(c, http.StatusInternalServerError, "Retrieving Data Failed")
	}

	return responseMultiplePost(c, http.StatusOK, "Success", data)
}

// GetPostById implements PostHandler.
func (p PostHttpHandler) GetPostById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response(c, http.StatusInternalServerError, "Retrieving ID Failed")
	}

	data, err := p.postUsecase.GetPostById(id)
	if err != nil {
		return response(c, http.StatusInternalServerError, "Retrieving Data Failed")
	}

	return responseMultiplePost(c, http.StatusOK, "Success", data)
}
