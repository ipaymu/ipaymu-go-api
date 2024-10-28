
# iPaymu Go Plugin

Golang (Go) package for iPaymu.com Indonesia Payment Gateway

## Documentation

For the API documentation, check [iPaymu API Reference](https://ipaymu.com/api-collection).

For the details of this library, see the [GoDoc](soon).

## Installation

Install ipaymu-go with:

```sh
go get github.com/ipaymu/ipaymu-go-api
```

Then, import it using:

```go
import (
    "github.com/ipaymu/ipaymu-go-api"
)
```

## Environment

## How to
1. First initiate the iPaymu `Client` (`ipaymu.Client`) that can be initiate with `NewClient()` function
2. With `Client` we can call api for payment (redirect, direct)
3. Each function for calling api payment have spesific type of request (`RequestRedirect`, `RequestDirectVA`, `RequestDirectConStore`, `RequestDirectCOD`) which have each constructor function

## Example
we have attached usage example in this repository folder `example/payment`
```go
func DirectVirtualAccount() error {
	client := ipaymu.NewClient()
	client.EnvApi = ipaymu.Sandbox
	client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
	client.VirtualAccount = "1179000899"

	request := ipaymu.NewRequestDirectVA(ipaymu.CimbNiaga)
	request.AddBuyer("buyer", "phone", "email@test.com")
	request.NotifyUrl = "http://localhost/notify-url"
	request.Expired = 24
	request.ExpiredType = ipaymu.Hours
	request.ReferenceId = "trx-123"
	request.Amount = 100000

	va, err := client.DirectPaymentVA(*request)
	if err != nil {
		return err
	}

	fmt.Println(va)

	return nil
}
```


## License

See [License](https://choosealicense.com/licenses/mit/)


## Support

For support, email us at support@ipaymu.com.

