package requests

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"websocket/handlers"
)

type SendMessageRequest struct {
	Message        string `json:"message" validate:"required"`
	UserIdentifier string `json:"user_identifier" validate:"required"`
}

var validate *validator.Validate

func Validate(w http.ResponseWriter, d SendMessageRequest) {
	validate = validator.New()

	err := validate.Struct(d)
	if err != nil {
		handlers.Throw400(w, err)
		return
	}
}
