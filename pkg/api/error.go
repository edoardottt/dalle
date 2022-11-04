package api

import "fmt"

type ErrorResponse struct {
	ErrorData Error `json:"error,omitempty"`
}

type Error struct {
	Message string      `json:"message,omitempty"`
	Type    string      `json:"type,omitempty"`
	Param   string      `json:"param,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s %s: %s", e.ErrorData.Type, e.ErrorData.Code, e.ErrorData.Message)
}
