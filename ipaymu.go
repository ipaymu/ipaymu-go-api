package ipaymu_go_api

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ClientApi interface {
	CallApi(url *url.URL, signature string, body []byte) ([]byte, error)
	CheckTransaction(transactionID int) (res ResponseCheck, err error)
	HistoryTransaction(request RequestTransactionHistory) (res ResponseTransaction, err error)
	ListPaymentMethod() (res ResponseListPayment, err error)
	DirectPaymentVA(request RequestDirectVA) (res Response, err error)
	DirectPaymentConStore(request RequestDirectConStore) (res Response, err error)
	DirectPaymentCOD(request RequestDirectCOD) (res Response, err error)
	RedirectPayment(request RequestRedirect) (res Response, err error)
	GetBalance() (res ResponseBalance, err error)
	AssignCredential(apiKey, virtualAccount string, env EnvironmentType)
}

type Client struct {
	ApiKey         string
	VirtualAccount string
	EnvApi         EnvironmentType
}

func NewClient() *Client {
	return &Client{
		EnvApi: Production,
	}
}

// AssignCredential sets the API key, virtual account, and environment for the iPaymu client.
//
// apiKey: The API key provided by iPaymu. This is required for authentication and authorization.
// virtualAccount: The virtual account number associated with the API key. This is used to identify the merchant account.
// env: The environment type (Production or Sandbox) to which the client will connect.
func (c *Client) AssignCredential(apiKey, virtualAccount string, env EnvironmentType) {
	c.ApiKey = apiKey
	c.VirtualAccount = virtualAccount
	c.EnvApi = env
}

var defHTTPTimeout = 30 * time.Second

// CallApi sends a POST request to the specified URL with the provided signature and body.
// It constructs an HTTP request with the necessary headers and makes a request to the iPaymu API.
//
// url: The URL to which the request will be sent.
// signature: The signature generated using the API key and other relevant data.
// body: The request body in JSON format.
//
// The function returns a byte slice containing the response body and an error if any occurred during the request.
func (c *Client) CallApi(url *url.URL, signature string, body []byte) ([]byte, error) {
	reqBody := io.NopCloser(strings.NewReader(string(body)))

	req := &http.Request{
		Method: http.MethodPost,
		URL:    url,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
			"va":           {c.VirtualAccount},
			"signature":    {signature},
			"Accept":       {"application/json"},
		},
		Body: reqBody,
	}

	httpClient := http.DefaultClient
	httpClient.Timeout = defHTTPTimeout
	resp, err := httpClient.Do(req)

	if err != nil {
		log.Printf("An Error Occured %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return respBody, nil
}
