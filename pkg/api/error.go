package api

import "fmt"

type Error struct {
	Message string      `json:"message,omitempty"`
	Type    string      `json:"type,omitempty"`
	Param   string      `json:"param,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Type, e.Code, e.Message)
}
