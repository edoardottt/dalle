package api

type Choices struct {
	Text         string      `json:"text,omitempty"`
	Index        int         `json:"index,omitempty"`
	Logprobs     interface{} `json:"logprobs,omitempty"`
	FinishReason string      `json:"finish_reason,omitempty"`
}
