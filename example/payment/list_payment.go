package payment

import (
	"fmt"
	ipaymu "github.com/ipaymu/ipaymu-go-api"
)

// ListPaymentMethod retrieves a list of available payment methods from iPaymu API.
//
// This function initializes a new iPaymu client with the sandbox environment, API key, and virtual account.
// It then calls the ListPaymentMethod method of the client to retrieve the list of payment methods.
// If an error occurs during the API call, it returns the error. Otherwise, it prints the retrieved payment methods and returns nil.
func ListPaymentMethod() error {
    // initiate client
    client := ipaymu.NewClient()
    client.EnvApi = ipaymu.Sandbox
    client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
    client.VirtualAccount = "1179000899"

    // api call
    balance, err := client.ListPaymentMethod()
    if err != nil {
        return err
    }

    fmt.Printf("%v", balance)

    return nil
}
