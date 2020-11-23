package models

type PostTimeout struct {
	Timeout int `json:"timeout"`
}

type Response struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
