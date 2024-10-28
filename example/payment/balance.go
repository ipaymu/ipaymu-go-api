package payment

import (
	"fmt"
	ipaymu "github.com/ipaymu/ipaymu-go-api"
)

// Balance retrieves the current balance of the specified virtual account.
//
// This function initiates a client with the provided sandbox environment, API key, and virtual account.
// It then calls the GetBalance method of the client to retrieve the balance.
// If an error occurs during the API call, it is returned. Otherwise, the balance is printed and nil is returned.
func Balance() error {
    // initiate client
    client := ipaymu.NewClient()
    client.EnvApi = ipaymu.Sandbox
    client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
    client.VirtualAccount = "1179000899"

    // api call
    balance, err := client.GetBalance()
    if err != nil {
        return err
    }

    fmt.Printf("%v", balance)

    return nil
}
