package ipaymu_go_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// GenerateSignature generates a signature for iPaymu API requests.
//
// The function takes three parameters:
// - body: A string representing the request body.
// - method: A string representing the HTTP method (e.g., "POST").
// - cfg: A Client struct containing configuration details for the iPaymu API.
//
// The function returns a string representing the generated signature.
//
// The signature is generated using the HMAC-SHA256 algorithm with the API key as the secret key.
// The input string to HMAC is constructed as follows:
// "POST:<virtual_account>:<lowercase_body_hash>:<api_key>"
//
// The body hash is calculated as the SHA256 hash of the request body.
// The lowercase body hash is then converted to a hexadecimal string.
//
// The generated signature is then returned as a hexadecimal string.
func GenerateSignature(body string, method string, cfg Client) string {
    bodyHash := sha256.Sum256([]byte(body))
    bodyHashToString := hex.EncodeToString(bodyHash[:])
    stringToSign := "POST:" + cfg.VirtualAccount + ":" + strings.ToLower(bodyHashToString) + ":" + cfg.ApiKey

    h := hmac.New(sha256.New, []byte(cfg.ApiKey))
    h.Write([]byte(stringToSign))
    return hex.EncodeToString(h.Sum(nil))
}
