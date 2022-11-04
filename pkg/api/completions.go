package api

type Completions struct {
	ID      string    `json:"id,omitempty"`
	Object  string    `json:"object,omitempty"`
	Created int       `json:"created,omitempty"`
	Model   string    `json:"model,omitempty"`
	Choices []Choices `json:"choices,omitempty"`
	Usage   Usage     `json:"usage,omitempty"`
}
