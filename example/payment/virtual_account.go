package payment

import (
	"fmt"
	ipaymu "github.com/ipaymu/ipaymu-go-api"
	"time"
)

// DirectVirtualAccount is a function that performs a direct payment to a virtual account using iPaymu API.
// It initiates a client with the sandbox environment, sets the API key and virtual account, prepares a request
// with buyer details, notification URL, expiration, and reference ID, and then makes an API call to process the payment.
//
// Parameters:
// None
//
// Return:
// error: An error object if the API call fails or any other error occurs during the process.
//        If the API call is successful, it returns nil.
func DirectVirtualAccount() error {
    // initiate client
    client := ipaymu.NewClient()
    client.EnvApi = ipaymu.Sandbox
    client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
    client.VirtualAccount = "1179000899"

    // prepare the request
    var exp int8 = 24
    var expType ipaymu.ExpiredType = ipaymu.Hours
    var refId string = time.Now().Format("20060102150405") // change based on needs
    var notifUrl string = "http://localhost/notify-url"
    request := ipaymu.NewRequestDirectVA(ipaymu.CimbNiaga)
    request.AddBuyer("buyer", "08123456789", "email@test.com")
    request.NotifyUrl = &notifUrl
    request.Expired = &exp
    request.ExpiredType = &expType
    request.ReferenceId = &refId
    request.Amount = 100000

    // api call
    va, err := client.DirectPaymentVA(*request)
    if err != nil {
        return err
    }

    fmt.Println(*va.Data)

    return nil
}
