package api

type Completions struct {
	ID      string    `json:"id,omitempty"`
	Object  string    `json:"object,omitempty"`
	Created int       `json:"created,omitempty"`
	Model   string    `json:"model,omitempty"`
	Choices []Choices `json:"choices,omitempty"`
	Usage   Usage     `json:"usage,omitempty"`
}

type CompletionsInput struct {
	Model            string      `json:"model,omitempty"`
	Prompt           string      `json:"prompt,omitempty"`
	Suffix           string      `json:"suffix,omitempty"`
	MaxTokens        int         `json:"max_tokens,omitempty"`
	Temperature      int         `json:"temperature,omitempty"`
	TopP             int         `json:"top_p,omitempty"`
	N                int         `json:"n,omitempty"`
	Stream           bool        `json:"stream,omitempty"`
	Logprobs         interface{} `json:"logprobs,omitempty"`
	Echo             bool        `json:"echo,omitempty"`
	Stop             string      `json:"stop,omitempty"`
	PresencePenalty  int         `json:"presence_penalty,omitempty"`
	FrequencePenalty int         `json:"frequence_penalty,omitempty"`
	BestOf           int         `json:"best_of,omitempty"`
	LogitBias        map[int]int `json:"logit_bias,omitempty"`
	User             string      `json:"user,omitempty"`
}
