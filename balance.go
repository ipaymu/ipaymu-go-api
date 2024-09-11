package ipaymu_go_api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

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
