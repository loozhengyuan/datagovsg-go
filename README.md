# Data.gov.sg

[![PkgGoDev](https://pkg.go.dev/badge/github.com/loozhengyuan/datagovsg-go/datagovsg)](https://pkg.go.dev/github.com/loozhengyuan/datagovsg-go/datagovsg)
[![Go Report Card](https://goreportcard.com/badge/github.com/loozhengyuan/datagovsg-go)](https://goreportcard.com/report/github.com/loozhengyuan/datagovsg-go)
![ci](https://github.com/loozhengyuan/datagovsg-go/workflows/ci/badge.svg)

Go wrapper for Data.gov.sg real-time [APIs](https://data.gov.sg/developer).

**Datasets**

The list of supported datasets are as follows:

|Category|Dataset|Endpoint|Supported|
|---|---|---|---|
|Transport|[Traffic Images](https://data.gov.sg/dataset/traffic-images?resource_id=e127e29a-bd48-47e2-a0a7-e89ce31f10c7)|`/v1/transport/traffic-images`|✅|
|Transport|[Taxi Availability](https://data.gov.sg/dataset/taxi-availability?resource_id=9d217820-1350-4032-a7a3-3cd83e222eb7)|`/v1/transport/taxi-availability`|✅|
|Transport|[Carpark Availability](https://data.gov.sg/dataset/carpark-availability?resource_id=4f4a57d1-e904-4326-b83e-dae99358edf9)|`/v1/transport/carpark-availability`|✅|
|Environment|[PM2.5](https://data.gov.sg/dataset/pm2-5?resource_id=fa0958a9-bade-419e-9475-cbf5ccf4f746)|`/v1/environment/pm25`|✅|
|Environment|[PSI](https://data.gov.sg/dataset/psi?resource_id=82776919-0de1-4faf-bd9e-9c997f9a729d)|`/v1/environment/psi`|✅|
|Environment|[Ultra-violet Index](https://data.gov.sg/dataset/ultraviolet-index-uvi?resource_id=6246c980-21d4-441f-a1d0-b321e2085420)|`/v1/environment/uv-index`|✅|
|Environment|[Air Temperature](https://data.gov.sg/dataset/realtime-weather-readings?resource_id=17494bed-23e9-4b3b-ae89-232f87987163)|`/v1/environment/air-temperature`|✅|
|Environment|[Rainfall](https://data.gov.sg/dataset/realtime-weather-readings?resource_id=8bd37e06-cdd7-4ca4-9ad8-5754eb70a33d)|`/v1/environment/rainfall`|✅|
|Environment|[Relative Humidity](https://data.gov.sg/dataset/realtime-weather-readings?resource_id=59eb2883-2ceb-4d16-85f0-7e3a3176ef46)|`/v1/environment/relative-humidity`|✅|
|Environment|[Wind Direction](https://data.gov.sg/dataset/realtime-weather-readings?resource_id=5dcf8aa5-cf6a-44e4-b359-1173eecfdf4c)|`/v1/environment/wind-direction`|✅|
|Environment|[Wind Speed](https://data.gov.sg/dataset/realtime-weather-readings?resource_id=16035f22-37b4-4a5c-b024-ca2381f11b48)|`/v1/environment/wind-speed`|✅|
|Environment|[2-hour Weather Forecast](https://data.gov.sg/dataset/weather-forecast?resource_id=571ef5fb-ed31-48b2-85c9-61677de42ca9)|`/v1/environment/2-hour-weather-forecast`|✅|
|Environment|[24-hour Weather Forecast](https://data.gov.sg/dataset/weather-forecast?resource_id=9a8bd97e-0e38-46b7-bc39-9a2cb4a53a62)|`/v1/environment/24-hour-weather-forecast`|✅|
|Environment|[4-day Weather Forecast](https://data.gov.sg/dataset/weather-forecast?resource_id=4df6d890-f23e-47f0-add1-fd6d580447d1)|`/v1/environment/4-day-weather-forecast`|✅|
|Technology|[IPOS Design Applications](https://data.gov.sg/dataset/ipos-apis?resource_id=adf6222f-955b-4a76-892f-802a396844a1)|`/v1/technology/ipos/designs`|🚧|
|Technology|[IPOS Trademark Applications](https://data.gov.sg/dataset/ipos-apis?resource_id=1522db0e-808b-48ea-9869-fe5adc566585)|`/v1/technology/ipos/trademarks`|🚧|
|Technology|[IPOS Patent Applications](https://data.gov.sg/dataset/ipos-apis?resource_id=6a030bf2-22da-4621-8ab0-9a5956a30ef3)|`/v1/technology/ipos/patents`|🚧|

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

For all supported APIs, the wrapper method is accessible using the `GetXxx` convention. For example, to fetch the latest traffic images, one would call the `GetTrafficImages()` method.

### Fetching the latest data

To illustrate this example, let us assume that we intend fetch the latest [traffic images](https://data.gov.sg/dataset/traffic-images) from the API.

```go
package main

import (
	"fmt"

	"github.com/loozhengyuan/datagovsg-go/datagovsg"
)

func main() {
	// Create api client
	c := datagovsg.NewClient()

	// Fetch latest traffic images
	img, _ := c.GetTrafficImages()
	for _, camera := range img.Items {
		for _, images := range camera.Cameras {
			fmt.Println(images.Image)
		}
	}
}
```

### Using the `datagovsg.QueryOption`

Most APIs allow you to pass a `date_time` or `date` parameter to retrieve data at a certain point in time. To do this, one can pass the `datagovsg.QueryOption` as a variadic argument when calling the respective API methods. Note that the client does not validate these options and will pass them to the API directly.

Using the above example, we will modify the above code to return the traffic images at `2020-05-01T08:03:00` using the `date_time` parameter:

```go
package main

import (
	"fmt"

	"github.com/loozhengyuan/datagovsg-go/datagovsg"
)

func main() {
	// Create api client
	c := datagovsg.NewClient()

	// Fetch traffic images at point in time
	// using datagovsg.QueryOption
	img, _ := c.GetTrafficImages(
		&datagovsg.QueryOption{
			Key:   "date_time",
			Value: "2020-05-01T08:03:00",
		},
	)
	for _, camera := range img.Items {
		for _, images := range camera.Cameras {
			fmt.Println(images.Image)
		}
	}
}
```

## License

[GPL-3.0](https://choosealicense.com/licenses/gpl-3.0/)

All data, including those used as test fixtures and those that you intend to retrieve using this package, are licensed under the [Singapore Open Data License](https://data.gov.sg/open-data-licence). All other aspects of the package is licensed under [GPL-3.0](https://choosealicense.com/licenses/gpl-3.0/).
