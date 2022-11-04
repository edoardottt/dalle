package api

type Edit struct {
	Object  string    `json:"object,omitempty"`
	Created int64     `json:"created,omitempty"`
	Choices []Choices `json:"choices,omitempty"`
	Usage   Usage     `json:"usage,omitempty"`
}
