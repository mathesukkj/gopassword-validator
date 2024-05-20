package models

type ResponseBody struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
