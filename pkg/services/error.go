package services

import (
	"github.com/edoardottt/dalle/pkg/api"
)

type ErrorResponse struct {
	Error api.Error `json:"error,omitempty"`
}
