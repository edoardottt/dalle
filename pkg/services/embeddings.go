package services

import (
	"context"

	"github.com/edoardottt/dalle/pkg/api"
)

// CreateEmbeddings creates an embedding vector representing the input text.
func (a *API) CreateEmbeddings(ctx context.Context, input *api.EmbeddingInput) (api.Embeddings, error) {
	var response api.Embeddings

	path := `/embeddings`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.Embeddings{}, err
	}

	return response, nil
}
