package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Error *[]string   `json:"error"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}, errors []string) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta:  meta,
		Data:  data,
		Error: &errors,
	}

	return jsonResponse
}

func ErrorResponse(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}