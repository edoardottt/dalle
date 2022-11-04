package main

import (
	"context"
	"fmt"

	"github.com/edoardottt/dalle"
)

func main() {
	d := dalle.New("API-KEY")

	models, err := d.Services.ListModels(context.TODO())

	if err != nil {
		fmt.Println(err)
	}

	for _, model := range models.Data {
		fmt.Println(model.ID)
	}
}
