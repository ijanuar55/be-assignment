package entity

type WebResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
