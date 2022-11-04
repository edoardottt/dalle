package services

import (
	"context"

	"github.com/edoardottt/dalle/pkg/api"
)

// CreateCompletion creates a completion for the provided prompt and parameters.
func (a *API) CreateCompletion(ctx context.Context, input *api.CompletionsInput) (api.Completions, error) {
	var response api.Completions

	path := `/completions`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.Completions{}, err
	}

	return response, nil
}
