package domain

type Response struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}
