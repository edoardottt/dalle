package api

type Moderations struct {
	ID      string    `json:"id,omitempty"`
	Model   string    `json:"model,omitempty"`
	Results []Results `json:"results,omitempty"`
}

type Categories struct {
	Hate            bool `json:"hate,omitempty"`
	HateThreatening bool `json:"hate/threatening,omitempty"`
	SelfHarm        bool `json:"self-harm,omitempty"`
	Sexual          bool `json:"sexual,omitempty"`
	SexualMinors    bool `json:"sexual/minors,omitempty"`
	Violence        bool `json:"violence,omitempty"`
	ViolenceGraphic bool `json:"violence/graphic,omitempty"`
}

type CategoryScores struct {
	Hate            float64 `json:"hate,omitempty"`
	HateThreatening float64 `json:"hate/threatening,omitempty"`
	SelfHarm        float64 `json:"self-harm,omitempty"`
	Sexual          float64 `json:"sexual,omitempty"`
	SexualMinors    float64 `json:"sexual/minors,omitempty"`
	Violence        float64 `json:"violence,omitempty"`
	ViolenceGraphic float64 `json:"violence/graphic,omitempty"`
}

type Results struct {
	Categories     Categories     `json:"categories,omitempty"`
	CategoryScores CategoryScores `json:"category_scores,omitempty"`
	Flagged        bool           `json:"flagged,omitempty"`
}

type ModerationInput struct {
	Model string `json:"model,omitempty"`
	Input string `json:"input,omitempty"`
}
