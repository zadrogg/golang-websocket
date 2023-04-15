package Validator

import (
	"errors"
	"net/http"
	"websocket/handlers"
)

func RuleMethod(w http.ResponseWriter, r *http.Request, m string) {
	if r.Method != m {
		err := errors.New("Method Not Allowed " + r.Method)
		handlers.Throw405(w, err)
	}
}
