package services

import (
	"context"
	"fmt"

	"github.com/edoardottt/dalle/pkg/api"
)

// ListFiles returns a list of files that belong to the user's organization.
func (a *API) ListFiles(ctx context.Context) (api.Files, error) {
	response := api.Files{}
	path := `/files`

	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.Files{}, err
	}

	return response, nil
}

// UploadFile uploads a file that contains document(s) to be used across various
// endpoints/features. Currently, the size of all the files uploaded by one organization
// can be up to 1 GB. Please contact us if you need to increase the storage limit.
func (a *API) UploadFile(ctx context.Context, input api.FileInput) (api.File, error) {
	var response api.File

	path := `files`
	if err := a.client.Post(ctx, path, &response, input); err != nil {
		return api.File{}, err
	}

	return response, nil
}

// DeleteFile deletes a file.
func (a *API) DeleteFile(ctx context.Context, fileID string) (api.File, error) {
	var response api.File

	path := fmt.Sprintf(
		`/files/%s`,
		fileID,
	)
	if err := a.client.Delete(ctx, path, &response); err != nil {
		return api.File{}, err
	}

	return response, nil
}

// RetrieveFile returns information about a specific file.
func (a *API) RetrieveFile(ctx context.Context, fileID string) (api.File, error) {
	var response api.File

	path := fmt.Sprintf(
		`/files/%s`,
		fileID,
	)
	if err := a.client.Get(ctx, path, &response); err != nil {
		return api.File{}, err
	}

	return response, nil
}

// RetrieveFile returns the contents of the specified file.
func (a *API) RetrieveFileContent(ctx context.Context, fileID string) ([]byte, error) {
	path := fmt.Sprintf(
		`/files/%s/content`,
		fileID,
	)
	resp, err := a.client.RawGet(ctx, path)

	if err != nil {
		return []byte{}, err
	}

	return resp, nil
}
