package api

type Images struct {
	Created int     `json:"created,omitempty"`
	Data    []Image `json:"data,omitempty"`
}

type Image struct {
	URL string `json:"url,omitempty"`
}
