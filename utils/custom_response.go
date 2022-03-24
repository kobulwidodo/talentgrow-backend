package utils

type Response struct {
	Message   string      `json:"message"`
	IsSuccess bool        `json:"is_success"`
	Data      interface{} `json:"data"`
}

func NewSuccessResponse(msg string, data interface{}) Response {
	return Response{
		Message:   msg,
		Data:      data,
		IsSuccess: true,
	}
}

func NewFailResponse(msg string) Response {
	return Response{
		Message:   msg,
		Data:      nil,
		IsSuccess: false,
	}
}
