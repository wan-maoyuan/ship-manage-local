package proto

import "net/http"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type PageList struct {
	List  any   `json:"list"`
	Count int64 `json:"count"`
}

func NewSuccesResponse(data any) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "请求成功",
		Data:    data,
	}
}

func NewFailedResponse(code int, msg string) *Response {
	return &Response{
		Code:    code,
		Message: msg,
	}
}
