package api

type Models struct {
	Object string  `json:"object,omitempty"`
	Data   []Model `json:"data,omitempty"`
}

type Model struct {
	ID         string       `json:"id,omitempty"`
	Object     string       `json:"object,omitempty"`
	Created    int          `json:"created,omitempty"`
	OwnedBy    string       `json:"owned_by,omitempty"`
	Permission []Permission `json:"permission,omitempty"`
	Root       string       `json:"root,omitempty"`
	Parent     interface{}  `json:"parent,omitempty"`
}
