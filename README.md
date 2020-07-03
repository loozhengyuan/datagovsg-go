# Data.gov.sg

![Main Workflow](https://github.com/loozhengyuan/datagovsg-go/workflows/Main%20Workflow/badge.svg)

Go wrapper for Data.gov.sg real-time [APIs](https://data.gov.sg/developer).

**Supported Datasets**

The list of supported datasets are as follows:

|Category|Dataset|Endpoint|Supported|
|---|---|---|---|
|Transport|[Traffic Images](https://data.gov.sg/dataset/traffic-images?resource_id=e127e29a-bd48-47e2-a0a7-e89ce31f10c7)|`/v1/transport/traffic-images`|✅|
|Transport|[Taxi Availability](https://data.gov.sg/dataset/taxi-availability?resource_id=9d217820-1350-4032-a7a3-3cd83e222eb7)|`/v1/transport/taxi-availability`|✅|
|Transport|[Carpark Availability](https://data.gov.sg/dataset/carpark-availability?resource_id=4f4a57d1-e904-4326-b83e-dae99358edf9)|`/v1/transport/carpark-availability`|✅|
|Environment|[PM2.5](https://data.gov.sg/dataset/pm2-5?resource_id=fa0958a9-bade-419e-9475-cbf5ccf4f746)|`/v1/environment/pm25`|✅|
|Environment|[PSI](https://data.gov.sg/dataset/psi?resource_id=82776919-0de1-4faf-bd9e-9c997f9a729d)|`/v1/environment/psi`|✅|

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
