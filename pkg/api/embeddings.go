package api

type Embeddings struct {
	Object string      `json:"object,omitempty"`
	Data   []Embedding `json:"data,omitempty"`
	Model  string      `json:"model,omitempty"`
	Usage  Usage       `json:"usage,omitempty"`
}

type Embedding struct {
	Object    string    `json:"object,omitempty"`
	Index     int       `json:"index,omitempty"`
	Embedding []float64 `json:"embedding,omitempty"`
}

type EmbeddingInput struct {
	Model string `json:"model,omitempty"`
	Input string `json:"input,omitempty"`
	User  string `json:"user,omitempty"`
}
