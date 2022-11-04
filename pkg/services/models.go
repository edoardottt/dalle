package services

import (
	"context"
	"fmt"

	"github.com/edoardottt/dalle/pkg/api"
)

// RetrieveModel retrieves a model instance, providing basic information
// about the model such as the owner and permissioning.
func (a *API) RetrieveModel(ctx context.Context, model string) (api.Model, error) {
	response := api.Model{}
	path := fmt.Sprintf(
		`/models/%s`,
		model,
	)

	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.Model{}, err
	}

	return response, nil
}

// ListModels lists the currently available models, and provides basic information
// about each one such as the owner and availability.
func (a *API) ListModels(ctx context.Context) (api.Models, error) {
	response := api.Models{}
	path := `/models`

	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.Models{}, err
	}

	return response, nil
}
