package client

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/karim-w/stdlib/httpclient"
)

type Client struct {
	key string
}

func New(
	key string,
) *Client {
	return &Client{key}
}

func createAuthHeader(
	method string,
	host string,
	resourcePath string,
	date string,
	secret string,
	body []byte,
) (string, string) {
	contentHash := computeContentHash(body)
	stringToSign := fmt.Sprintf("%s\n%s\n%s;%s;%s", method, resourcePath, date, host, contentHash)
	signature := computeSignature(stringToSign, secret)

	// Concatenate the string, which will be used in the authorization header.
	authorizationHeader := fmt.Sprintf(
		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
		signature,
	)

	return contentHash, authorizationHeader
}

func computeContentHash(content []byte) string {
	sha256 := sha256.New()
	sha256.Write(content)
	hashedBytes := sha256.Sum(nil)
	base64EncodedBytes := base64.StdEncoding.EncodeToString(hashedBytes)
	return base64EncodedBytes
}

func computeSignature(stringToSign string, secret string) string {
	decodedSecret, _ := base64.StdEncoding.DecodeString(secret)
	hash := hmac.New(sha256.New, decodedSecret)
	hash.Write([]byte(stringToSign))
	hashedBytes := hash.Sum(nil)
	encodedSignature := base64.StdEncoding.EncodeToString(hashedBytes)
	return encodedSignature
}

func (c *Client) Patch(
	ctx context.Context,
	host string,
	resource string,
	query string,
	reqbody interface{},
	response interface{},
) error {
	body := []byte("{}")
	var err error
	if reqbody != nil {
		body, err = json.Marshal(reqbody)
		if err != nil {
			return err
		}
	} else {
		reqbody = struct{}{}
	}
	date := time.Now().UTC().Format(http.TimeFormat)
	// Compute a content hash for the 'x-ms-content-sha256' header.
	contentHash := computeContentHash(body)
	// Prepare a string to sign.
	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", resource+"?"+query, date, host, contentHash)

	// Compute the signature.
	signature := computeSignature(stringToSign, c.key)

	// Concatenate the string, which will be used in the authorization header.
	authorizationHeader := fmt.Sprintf(
		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
		signature,
	)
	res := httpclient.Req(
		"https://"+host+resource+"?"+query,
	).AddHeader(
		"x-ms-date", date,
	).AddHeader(
		"x-ms-content-sha256", contentHash,
	).AddHeader(
		"Authorization", authorizationHeader,
	).AddHeader(
		"Content-Type", "application/json",
	).AddBody(
		reqbody,
	).Post()
	responseBody := res.GetBody()
	if !res.IsSuccess() {
		return fmt.Errorf("request failed: %s", string(responseBody))
	}
	if len(responseBody) == 0 {
		return nil
	}
	return json.Unmarshal(responseBody, &response)
}

func (c *Client) Post(
	ctx context.Context,
	host string,
	resource string,
	query string,
	reqbody interface{},
	response interface{},
) error {
	body := []byte("{}")
	var err error
	if reqbody != nil {
		body, err = json.Marshal(reqbody)
		if err != nil {
			return err
		}
	} else {
		reqbody = struct{}{}
	}
	date := time.Now().UTC().Format(http.TimeFormat)
	// Compute a content hash for the 'x-ms-content-sha256' header.
	contentHash := computeContentHash(body)
	// Prepare a string to sign.
	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", resource+"?"+query, date, host, contentHash)

	// Compute the signature.
	signature := computeSignature(stringToSign, c.key)

	// Concatenate the string, which will be used in the authorization header.
	authorizationHeader := fmt.Sprintf(
		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
		signature,
	)
	res := httpclient.Req(
		"https://"+host+resource+"?"+query,
	).AddHeader(
		"x-ms-date", date,
	).AddHeader(
		"x-ms-content-sha256", contentHash,
	).AddHeader(
		"Authorization", authorizationHeader,
	).AddHeader(
		"Content-Type", "application/json",
	).AddBody(
		reqbody,
	).Post()
	responseBody := res.GetBody()
	if !res.IsSuccess() {
		return fmt.Errorf("request failed: %s", string(responseBody))
	}
	if len(responseBody) == 0 {
		return nil
	}
	return json.Unmarshal(responseBody, &response)
}

func (c *Client) Delete(
	ctx context.Context,
	host string,
	resource string,
	query string,
	response interface{},
) error {
	body := []byte("{}")
	reqBody := struct{}{}
	date := time.Now().UTC().Format(http.TimeFormat)
	contentHash, authHeader := createAuthHeader(
		"DELETE",
		host,
		resource+"?"+query,
		date,
		c.key,
		body,
	)
	res := httpclient.Req(
		"https://"+host+resource+"?"+query,
	).AddHeader(
		"x-ms-date", date,
	).AddHeader(
		"x-ms-content-sha256", contentHash,
	).AddHeader(
		"Authorization", authHeader,
	).AddHeader(
		"Content-Type", "application/json",
	).AddBody(reqBody).Del()
	responseBody := res.GetBody()
	if !res.IsSuccess() {
		return fmt.Errorf("request failed: %s", string(responseBody))
	}
	if len(responseBody) == 0 {
		return nil
	}
	return json.Unmarshal(responseBody, &response)
}

func (c *Client) Get(
	ctx context.Context,
	host string,
	resource string,
	query string,
	response interface{},
) error {
	body := []byte("{}")
	reqBody := struct{}{}
	date := time.Now().UTC().Format(http.TimeFormat)
	contentHash, authHeader := createAuthHeader(
		"GET",
		host,
		resource+"?"+query,
		date,
		c.key,
		body,
	)
	res := httpclient.Req(
		"https://"+host+resource+"?"+query,
	).AddHeader(
		"x-ms-date", date,
	).AddHeader(
		"x-ms-content-sha256", contentHash,
	).AddHeader(
		"Authorization", authHeader,
	).AddHeader(
		"Content-Type", "application/json",
	).AddBody(reqBody).Get()
	responseBody := res.GetBody()
	if !res.IsSuccess() {
		return fmt.Errorf("request failed: %s", string(responseBody))
	}
	if len(responseBody) == 0 {
		return nil
	}
	return json.Unmarshal(responseBody, &response)
}
