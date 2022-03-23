package utils

type Response struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"is_success"`
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
