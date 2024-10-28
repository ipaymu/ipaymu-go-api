package ipaymu_go_api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetBalance retrieves the balance of the specified virtual account.
//
// The function sends a POST request to the iPaymu API endpoint for balance retrieval.
// It uses the provided Client instance, which contains the necessary configuration and authentication details.
//
// Parameters:
// - c: A pointer to the Client instance, which holds the API endpoint, virtual account, and authentication details.
//
// Return values:
// - res: A ResponseBalance struct containing the API response data.
// - err: An error if any occurred during the API call or response parsing.
//
// If the API call is successful and the status code is 200, the function returns the ResponseBalance struct.
// If the API call fails or the status code is not 200, the function returns an error with the corresponding message.

func (c *Client) GetBalance() (res ResponseBalance, err error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/balance", c.EnvApi))
	jsonBody, _ := json.Marshal(map[string]string{"account": c.VirtualAccount})
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
	api, err := c.CallApi(url, signature, jsonBody)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(api, &res)
	if err != nil {
		return res, err
	}

	if res.Status != 200 {
		return res, fmt.Errorf("%s\n", res.Message)
	}

	return
}
