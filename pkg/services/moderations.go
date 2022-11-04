package services

import (
	"context"

	"github.com/edoardottt/dalle/pkg/api"
)

// CreateModeration classifies if text violates OpenAI's Content Policy.
func (a *API) CreateModeration(ctx context.Context, input *api.ModerationInput) (api.Moderations, error) {
	var response api.Moderations

	path := `/moderations`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.Moderations{}, err
	}

	return response, nil
}
