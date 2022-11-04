package api

type Permission struct {
	ID                 string      `json:"id,omitempty"`
	Object             string      `json:"object,omitempty"`
	Created            int         `json:"created,omitempty"`
	AllowCreateEngine  bool        `json:"allow_create_engine,omitempty"`
	AllowSampling      bool        `json:"allow_sampling,omitempty"`
	AllowLogprobs      bool        `json:"allow_logprobs,omitempty"`
	AllowSearchIndices bool        `json:"allow_search_indices,omitempty"`
	AllowView          bool        `json:"allow_view,omitempty"`
	AllowFineTuning    bool        `json:"allow_fine_tuning,omitempty"`
	Organization       string      `json:"organization,omitempty"`
	Group              interface{} `json:"group,omitempty"`
	IsBlocking         bool        `json:"is_blocking,omitempty"`
}
