package ipaymu_go_api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// CheckTransaction is used to check the status of a specific transaction by its ID.
//
// transactionID: The unique identifier of the transaction to be checked.
//
// Returns:
// res: A ResponseCheck struct containing the response data from the API.
// err: An error if any occurred during the API call or response parsing.
//
// Note: If the API call is successful and the transaction status is not 200, an error will be returned with the corresponding message.
func (c *Client) CheckTransaction(transactionID int) (res ResponseCheck, err error) {
    uri, _ := url.Parse(fmt.Sprintf("%s/api/v2/transaction", c.EnvApi))
    jsonBody, _ := json.Marshal(map[string]int{"transactionId": transactionID})
    signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
    api, err := c.CallApi(uri, signature, jsonBody)
    if err != nil {
        return res, err
    }

    err = json.Unmarshal(api, &res)
    if err != nil {
        return res, err
    }

    if res.Status != 200 {
        return res, fmt.Errorf("%s", res.Message)
    }

    return
}

// HistoryTransaction retrieves transaction history based on the provided request parameters.
//
// Parameters:
// request: A RequestTransactionHistory struct containing the necessary parameters for the transaction history request.
//
// Returns:
// res: A ResponseTransaction struct containing the response data from the API.
// err: An error if any occurred during the API call or response parsing.
//
// Note: If the API call is successful and the transaction history status is not 200, an error will be returned with the corresponding message.
func (c *Client) HistoryTransaction(request RequestTransactionHistory) (res ResponseTransaction, err error) {
    uri, _ := url.Parse(fmt.Sprintf("%s/api/v2/history", c.EnvApi))
    jsonBody, _ := json.Marshal(request)
    signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
    api, err := c.CallApi(uri, signature, jsonBody)
    if err != nil {
        return
    }

    err = json.Unmarshal(api, &res)
    if err != nil {
        return
    }

    if res.Status != 200 {
        return res, fmt.Errorf("%s", res.Message)
    }

    return
}
