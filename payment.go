package ipaymu_go_api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c Client) ListPaymentMethod() (res ResponseListPayment, err error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment-method-list", c.EnvApi))
	jsonBody, _ := json.Marshal(map[string]bool{"request": true})
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	api, err := c.CallApi(url, signature, jsonBody)
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

func (c Client) DirectPaymentVA(request RequestDirectVA) (res Response, err error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
	jsonBody, _ := json.Marshal(request)
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	api, err := c.CallApi(url, signature, jsonBody)
	if err != nil {
		return Response{}, err
	}

	err = json.Unmarshal(api, &res)
	if err != nil {
		return Response{}, err
	}

	if res.Status != 200 {
		return res, fmt.Errorf("%s", res.Message)
	}

	return
}

func (c Client) DirectPaymentConStore(request RequestDirectConStore) (res Response, err error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
	jsonBody, _ := json.Marshal(request)
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	api, err := c.CallApi(url, signature, jsonBody)
	if err != nil {
		return Response{}, err
	}

	err = json.Unmarshal(api, &res)
	if err != nil {
		return Response{}, err
	}

	if res.Status != 200 {
		return res, fmt.Errorf("%s", res.Message)
	}

	return
}

func (c Client) DirectPaymentCOD(request RequestDirectCOD) (res Response, err error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
	jsonBody, _ := json.Marshal(request)
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	api, err := c.CallApi(url, signature, jsonBody)
	if err != nil {
		return Response{}, err
	}

	err = json.Unmarshal(api, &res)
	if err != nil {
		return Response{}, err
	}

	if res.Status != 200 {
		return res, fmt.Errorf("%s", res.Message)
	}

	return
}

func (c Client) RedirectPayment(request RequestRedirect) (res Response, err error) {
	url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/", c.EnvApi))
	jsonBody, _ := json.Marshal(request)
	signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", c))
	api, err := c.CallApi(url, signature, jsonBody)
	if err != nil {
		return Response{}, err
	}

	err = json.Unmarshal(api, &res)
	if err != nil {
		return Response{}, err
	}

	if res.Status != 200 {
		return res, fmt.Errorf("%s", res.Message)
	}

	return
}
