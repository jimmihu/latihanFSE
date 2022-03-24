package dto

import (
	"log"
	"net/http"
)

type HttpResponse struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Error      *string     `json:"error"`
	Data       interface{} `json:"data"`
}

func ErrorDBResponse(err error) HttpResponse {

	log.Print(err)
	errMsg := "Database Error"
	return HttpResponse{
		StatusCode: http.StatusInternalServerError,
		Status:     "failed",
		Error:      &errMsg,
		Data:       nil,
	}
}
