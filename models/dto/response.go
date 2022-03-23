package dto

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Error      *string     `json:"error"`
	Data       interface{} `json:"data"`
}
