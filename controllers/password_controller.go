package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/mathesukkj/gopassword-validator/models"
	"github.com/mathesukkj/gopassword-validator/services"
)

func ValidatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	validate := validator.New()

	body, err := io.ReadAll(r.Body)

	var payload models.ValidatePasswordPayload
	if err = json.Unmarshal([]byte(body), &payload); err != nil {
		ApiResponse(w, models.ResponseBody{
			Message: "error while reading request body!",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	if err := validate.Struct(payload); err != nil {
		ApiResponse(w, models.ResponseBody{
			Message: "wrong body format!",
			Error:   err.Error(),
		}, http.StatusBadRequest)
		return
	}

	errors := services.ValidatePassword(payload.Password)
	if len(errors) > 0 {
		ApiResponse(w, models.ResponseBody{
			Message: "password didnt pass the validation!",
			Error:   strings.Join(errors, ", "),
		}, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ApiResponse(w http.ResponseWriter, body models.ResponseBody, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&body)
}
