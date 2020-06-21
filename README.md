# Data.gov.sg

![Main Workflow](https://github.com/loozhengyuan/datagovsg-go/workflows/Main%20Workflow/badge.svg)

Go wrapper for Data.gov.sg real-time [APIs](https://data.gov.sg/developer).

## Installation

To install the wrapper, use `go get` to fetch the latest version:

```shell
go get -u github.com/loozhengyuan/datagovsg-go/datagovsg
```

Once installed, import the `datagovsg` package in your Go application:

```shell
import "github.com/loozhengyuan/datagovsg-go/datagovsg"
```

## Usage

```go
package main

import (
	"log"

	"github.com/loozhengyuan/datagovsg-go/datagovsg"
)

func main() {
	// Create api client
	c := datagovsg.NewClient()

	// Fetch latest traffic images
	img, err := c.GetTrafficImages()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	for _, camera := range img.Items {
		for _, images := range camera.Cameras {
			log.Println(images.Image)
		}
	}
}
```

## License

[GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
