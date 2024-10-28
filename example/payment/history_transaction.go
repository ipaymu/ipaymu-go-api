package payment

import (
	"fmt"
	ipaymu "github.com/ipaymu/ipaymu-go-api"
	"strings"
)

// HistoryTransaction retrieves transaction history from iPaymu API.
//
// This function initiates a client with the sandbox environment, sets the API key and virtual account,
// prepares a request for transaction history with specified status, order, and list of transaction IDs,
// and makes an API call to retrieve the transaction history.
//
// Parameters:
// - None
//
// Return:
// - error: An error object if the API call fails or if any other error occurs during the process.
//          If the API call is successful, it returns nil.
//
// Example usage:
//
//	err := HistoryTransaction()
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
func HistoryTransaction() error {
    // initiate client
    client := ipaymu.NewClient()
    client.EnvApi = ipaymu.Sandbox
    client.ApiKey = "QbGcoO0Qds9sQFDmY0MWg1Tq.xtuh1"
    client.VirtualAccount = "1179000899"

    // prepare request
    status := ipaymu.Success
    orderBy := ipaymu.Paid
    listID := []string{"68369", "44396", "44389"}
    listBulk := strings.Join(listID, ",")
    request := ipaymu.NewRequestTransactionHistory()
    request.Status = &status
    request.OrderBy = &orderBy
    request.BulkId = &listBulk

    // api call
    api, err := client.HistoryTransaction(*request)
    if err != nil {
        return err
    }

    fmt.Printf("%+v\n", api)
    return nil
}
