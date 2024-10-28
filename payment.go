package ipaymu_go_api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ListPaymentMethod retrieves a list of available payment methods from iPaymu API.
//
// The function sends a POST request to the iPaymu API endpoint "/api/v2/payment-method-list"
// with a JSON body containing a "request": true field. It then generates a signature
// using the provided client configuration and sends the request with the signature.
//
// If the request is successful (status code 200), the function unmarshals the response
// into a ResponseListPayment struct and returns it along with no error.
//
// If the request fails (status code other than 200), the function returns an error
// containing the error message from the response.
//
// If any error occurs during the request or response processing, the function returns
// an empty ResponseListPayment and the corresponding error.
func (c *Client) ListPaymentMethod() (res ResponseListPayment, err error) {
    url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment-method-list", c.EnvApi))
    jsonBody, _ := json.Marshal(map[string]bool{"request": true})
    signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
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

// DirectPaymentVA sends a direct payment request to iPaymu API using Virtual Account (VA) payment method.
//
// Parameters:
// - request: A RequestDirectVA struct containing the payment details such as customer information,
//            order details, and payment method specific details.
//
// Return:
// - res: A Response struct containing the API response status, message, and any additional data.
// - err: An error if any occurred during the API request or response processing.
//
// The function constructs the API endpoint URL, marshals the request into JSON, generates a signature,
// sends the request with the signature, and processes the response. If the request is successful (status code 200),
// it unmarshals the response into a Response struct and returns it along with no error. If the request fails
// (status code other than 200), it returns an error containing the error message from the response.
func (c *Client) DirectPaymentVA(request RequestDirectVA) (res Response, err error) {
    url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
    jsonBody, _ := json.Marshal(request)
    signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
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

// DirectPaymentConStore sends a direct payment request to iPaymu API using Conventional Store payment method.
//
// Parameters:
// - request: A RequestDirectConStore struct containing the payment details such as customer information,
//            order details, and payment method specific details.
//
// Return:
// - res: A Response struct containing the API response status, message, and any additional data.
// - err: An error if any occurred during the API request or response processing.
//
// The function constructs the API endpoint URL, marshals the request into JSON, generates a signature,
// sends the request with the signature, and processes the response. If the request is successful (status code 200),
// it unmarshals the response into a Response struct and returns it along with no error. If the request fails
// (status code other than 200), it returns an error containing the error message from the response.
func (c *Client) DirectPaymentConStore(request RequestDirectConStore) (res Response, err error) {
    url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
    jsonBody, _ := json.Marshal(request)
    signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
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

// DirectPaymentCOD sends a direct payment request to iPaymu API using Cash on Delivery (COD) payment method.
//
// Parameters:
// - request: A RequestDirectCOD struct containing the payment details such as customer information,
//            order details, and payment method specific details.
//
// Return:
// - res: A Response struct containing the API response status, message, and any additional data.
// - err: An error if any occurred during the API request or response processing.
//
// The function constructs the API endpoint URL, marshals the request into JSON, generates a signature,
// sends the request with the signature, and processes the response. If the request is successful (status code 200),
// it unmarshals the response into a Response struct and returns it along with no error. If the request fails
// (status code other than 200), it returns an error containing the error message from the response.
func (c *Client) DirectPaymentCOD(request RequestDirectCOD) (res Response, err error) {
    url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/direct", c.EnvApi))
    jsonBody, _ := json.Marshal(request)
    signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
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

// RedirectPayment sends a redirect payment request to iPaymu API.
//
// This function constructs a POST request to the iPaymu API endpoint "/api/v2/payment/"
// with a JSON body containing the payment details provided in the request parameter.
// It generates a signature using the provided client configuration and sends the request with the signature.
//
// Parameters:
// - request: A RequestRedirect struct containing the payment details such as customer information,
//            order details, and payment method specific details.
//
// Return:
// - res: A Response struct containing the API response status, message, and any additional data.
// - err: An error if any occurred during the API request or response processing.
//
// If the request is successful (status code 200), the function unmarshals the response
// into a Response struct and returns it along with no error.
//
// If the request fails (status code other than 200), the function returns an error
// containing the error message from the response.
//
// If any error occurs during the request or response processing, the function returns
// an empty Response and the corresponding error.
func (c *Client) RedirectPayment(request RequestRedirect) (res Response, err error) {
    url, _ := url.Parse(fmt.Sprintf("%s/api/v2/payment/", c.EnvApi))
    jsonBody, _ := json.Marshal(request)
    signature := fmt.Sprintf("%s", GenerateSignature(string(jsonBody), "POST", *c))
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
