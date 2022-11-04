package main

import (
	"context"
	"fmt"

	"github.com/edoardottt/dalle"
)

func main() {
	d := dalle.New("API-KEY")

	files, err := d.Services.ListFiles(context.TODO())

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files.Data {
		fmt.Println(file.ID)
	}
}
