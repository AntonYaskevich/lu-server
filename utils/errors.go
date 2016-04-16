package utils

type ApiError struct {
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Error       string `json:"error,omitempty"`
}
