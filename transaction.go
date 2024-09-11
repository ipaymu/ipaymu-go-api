package ipaymu_go_api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

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
