package api

type Image struct {
	Created int         `json:"created,omitempty"`
	Data    []ImageData `json:"data,omitempty"`
}

type ImageData struct {
	URL string `json:"url,omitempty"`
}

type ImageInput struct {
	Image          string `json:"image,omitempty"`
	Mask           int    `json:"mask,omitempty"`
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}
