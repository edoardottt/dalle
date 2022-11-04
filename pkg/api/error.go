package api

type Error struct {
	Message string      `json:"message,omitempty"`
	Type    string      `json:"type,omitempty"`
	Param   string      `json:"param,omitempty"`
	Code    interface{} `json:"code,omitempty"`
}
