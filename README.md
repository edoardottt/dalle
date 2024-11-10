# dalle

Simple Golang Client to interact with Dall-E API. See the [API documentation here](https://beta.openai.com/docs/api-reference).

Usage 🚀
-------

```Go
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

```

Contributing 🛠
-------

Just open an [issue](https://github.com/edoardottt/dalle/issues) / [pull request](https://github.com/edoardottt/dalle/pulls).

Before opening a pull request, download [golangci-lint](https://golangci-lint.run/usage/install/) and run

```bash
golangci-lint run
```

If there aren't errors, go ahead :)

Inspired by [liamg/hackerone](https://github.com/liamg/hackerone).

License 📝
-------

This repository is under [MIT License](https://github.com/edoardottt/dalle/blob/main/LICENSE).  
[edoardottt.com](https://edoardottt.com) to contact me.
