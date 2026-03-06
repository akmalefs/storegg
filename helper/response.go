package helper

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data any) Response {
	Meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := Response{
		Meta: Meta,
		Data: data,
	}

	return response
}
