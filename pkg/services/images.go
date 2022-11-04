package services

import (
	"context"

	"github.com/edoardottt/dalle/pkg/api"
)

// CreateImage creates an image given a prompt.
func (a *API) CreateImage(ctx context.Context, input *api.ImageInput) (api.Image, error) {
	var response api.Image

	path := `/images/generations`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.Image{}, err
	}

	return response, nil
}

// CreateImageEdit creates an edited or extended image given an original image and a prompt.
func (a *API) CreateImageEdit(ctx context.Context, input *api.ImageInput) (api.Image, error) {
	var response api.Image

	path := `/images/edits`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.Image{}, err
	}

	return response, nil
}

// CreateImageVariation creates a variation of a given image.
func (a *API) CreateImageVariation(ctx context.Context, input *api.ImageInput) (api.Image, error) {
	var response api.Image

	path := `/images/variations`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.Image{}, err
	}

	return response, nil
}
