package entity

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewAPIResponse(status int, message string, data interface{}) *APIResponse {
	return &APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
