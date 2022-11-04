package dalle

import (
	"github.com/edoardottt/dalle/pkg/client"
)

type API struct {
	Services *services.API
}

func New(apiKey string) *API {
	c := client.New(apiKey)

	return &API{
		Services: services.New(c),
	}
}
