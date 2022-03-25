package dto

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Error      *string     `json:"error"`
	Data       interface{} `json:"data"`
}

func DBErrorResponse(err error) HttpResponse {
	log.Print(err)
	errMsg := "Database Error"
	return HttpResponse{
		StatusCode: http.StatusInternalServerError,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}

func UserNotFoundResponse(err error) HttpResponse {
	log.Print(err)
	errMsg := "User Not Found"
	return HttpResponse{
		StatusCode: http.StatusNotFound,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}

func JsonRequestErrorResponse(c *gin.Context, err error) {
	log.Print(err)
	errMsg := "Invalid Json Request"
	c.JSON(http.StatusBadRequest, HttpResponse{
		StatusCode: http.StatusBadRequest,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	})
}
