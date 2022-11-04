package api

type Edit struct {
	Object  string    `json:"object,omitempty"`
	Created int64     `json:"created,omitempty"`
	Choices []Choices `json:"choices,omitempty"`
	Usage   Usage     `json:"usage,omitempty"`
}

type EditInput struct {
	Model       string `json:"model"`
	Input       string `json:"input"`
	Instruction string `json:"instruction"`
	N           int    `json:"n,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
	TopP        int    `json:"top_p,omitempty"`
}
