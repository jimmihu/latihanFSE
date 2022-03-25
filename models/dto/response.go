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
	errMsg := "Internal Server Error!"
	return HttpResponse{
		StatusCode: http.StatusInternalServerError,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}

func UserNotFoundResponse(err error) HttpResponse {
	log.Print(err)
	errMsg := "User Not Found!"
	return HttpResponse{
		StatusCode: http.StatusNotFound,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}

func ProductNotFoundResponse(err error) HttpResponse {
	log.Print(err)
	errMsg := "Product Not Found!"
	return HttpResponse{
		StatusCode: http.StatusNotFound,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}

func UnauthorizedResponse(err error) HttpResponse {
	log.Print(err)
	errMsg := "Unauthorized!"
	return HttpResponse{
		StatusCode: http.StatusUnauthorized,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}

func ForbiddenResponse(err error) HttpResponse {
	log.Print(err)
	errMsg := "Forbidden!"
	return HttpResponse{
		StatusCode: http.StatusForbidden,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}

func JsonRequestErrorResponse(c *gin.Context, err error) {
	log.Print(err)
	errMsg := "Bad Json Request!"
	c.JSON(http.StatusBadRequest, HttpResponse{
		StatusCode: http.StatusBadRequest,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	})
}
