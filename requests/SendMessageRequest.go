package requests

type SendMessageRequest struct {
	Message        string `json:"message"`
	UserIdentifier string `json:"user_identifier"`
}
