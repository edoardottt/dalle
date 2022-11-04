package api

type Files struct {
	Data   []File `json:"data,omitempty"`
	Object string `json:"object,omitempty"`
}

type File struct {
	ID        string `json:"id,omitempty"`
	Object    string `json:"object,omitempty"`
	Bytes     int    `json:"bytes,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	Filename  string `json:"filename,omitempty"`
	Purpose   string `json:"purpose,omitempty"`
	Deleted   bool   `json:"deleted,omitempty"`
}
