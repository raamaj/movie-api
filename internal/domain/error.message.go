package domain

type ErrorMessage struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
	Error      bool        `json:"error"`
}
