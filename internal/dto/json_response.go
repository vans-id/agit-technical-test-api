package dto

type JSONResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
