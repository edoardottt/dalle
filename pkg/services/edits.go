package services

import (
	"context"

	"github.com/edoardottt/dalle/pkg/api"
)

// CreateEdit creates a new edit for the provided input, instruction, and parameters.
func (a *API) CreateEdit(ctx context.Context, input *api.EditInput) (api.Edit, error) {
	var response api.Edit

	path := `/edits`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.Edit{}, err
	}

	return response, nil
}
