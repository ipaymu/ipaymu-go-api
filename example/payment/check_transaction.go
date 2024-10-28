package payment

import (
	"fmt"
	ipaymu "github.com/ipaymu/ipaymu-go-api"
)

// CheckTransaction is a function that checks the status of a transaction using iPaymu API.
// It initiates a client with the sandbox environment, API key, and virtual account,
// then calls the CheckTransaction method of the client with a transaction ID.
// If the API call is successful, it prints the transaction details and returns nil.
// If an error occurs during the API call, it returns the error.
func CheckTransaction() error {
    // initiate client
    client := ipaymu.NewClient()
    client.EnvApi = ipaymu.Sandbox
    client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
    client.VirtualAccount = "1179000899"

    // api call
    trx, err := client.CheckTransaction(96748)
    if err != nil {
        return err
    }

    fmt.Printf("%+v\n", trx)

    return nil
}
