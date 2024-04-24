package clientv2

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/karim-w/go-azure-communication-services/logger"
	"github.com/karim-w/stdlib/httpclient"
)

type client_ struct {
	key    string
	logger logger.Logger
}

func (c *client_) WithLogger(logger logger.Logger) Client {
	c.logger = logger
	return c
}

type Client interface {
	WithLogger(logger logger.Logger) Client
	Post(
		ctx context.Context,
		host string,
		resource string,
		queryParsms map[string][]string,
		reqbody any,
	) (io.Reader, error)
}

func (c *client_) Post(
	ctx context.Context,
	host string,
	resource string,
	queryParsms map[string][]string,
	reqbody any,
) (io.Reader, error) {
	body := []byte("{}")

	var err error
	if reqbody != nil {
		body, err = json.Marshal(reqbody)

		if err != nil {
			return nil, err
		}
	}

	resourceStringBuilder := strings.Builder{}

	resourceStringBuilder.WriteString(resource)

	if len(queryParsms) > 0 {
		resourceStringBuilder.WriteString("?")
	}

	for key, values := range queryParsms {
		for _, value := range values {
			resourceStringBuilder.WriteString(key)
			resourceStringBuilder.WriteString("=")
			resourceStringBuilder.WriteString(value)
			resourceStringBuilder.WriteString("&")
		}
	}

	resour := resourceStringBuilder.String()

	if len(resour) > 0 {
		resour = resour[:len(resour)-1]
	}

	date := time.Now().UTC().Format(http.TimeFormat)
	// Compute a content hash for the 'x-ms-content-sha256' header.
	contentHash := computeContentHash(body)
	// Prepare a string to sign.
	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", resour, date, host, contentHash)

	// Compute the signature.
	signature := computeSignature(stringToSign, c.key)

	// Concatenate the string, which will be used in the authorization header.
	authorizationHeader := fmt.Sprintf(
		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
		signature,
	)

	res := httpclient.Req(
		"https://"+host+resour,
	).AddHeader(
		"x-ms-date", date,
	).AddHeader(
		"x-ms-content-sha256", contentHash,
	).AddHeader(
		"Authorization", authorizationHeader,
	).AddHeader(
		"Content-Type", "application/json",
	).AddBodyRaw(
		body,
	).Post()

	responseBody := res.GetBody()

	if !res.IsSuccess() {
		code := res.GetStatusCode()

		c.logger.Error(
			"POST Request failed",
			logger.String("response", string(responseBody)),
			logger.Int("status_code", code),
			logger.String("CURL", res.CURL()),
		)

		return nil, errors.New(string(responseBody))
	}

	if len(responseBody) == 0 {
		return nil, nil
	}

	return bytes.NewReader(responseBody), nil
}

func New(
	key string,
) *client_ {
	return &client_{key, logger.Noop()}
}

func NewWithLogger(
	key string,
	logger logger.Logger,
) Client {
	return &client_{key, logger}
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

// func (c *client_) Patch(
// 	ctx context.Context,
// 	host string,
// 	resource string,
// 	query string,
// 	reqbody interface{},
// 	response interface{},
// ) error {
// 	body := []byte("{}")
// 	var err error
// 	if reqbody != nil {
// 		body, err = json.Marshal(reqbody)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		reqbody = struct{}{}
// 	}
// 	date := time.Now().UTC().Format(http.TimeFormat)
// 	// Compute a content hash for the 'x-ms-content-sha256' header.
// 	contentHash := computeContentHash(body)
// 	// Prepare a string to sign.
// 	stringToSign := fmt.Sprintf("PATCH\n%s\n%s;%s;%s", resource+"?"+query, date, host, contentHash)
//
// 	// Compute the signature.
// 	signature := computeSignature(stringToSign, c.key)
//
// 	// Concatenate the string, which will be used in the authorization header.
// 	authorizationHeader := fmt.Sprintf(
// 		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
// 		signature,
// 	)
// 	res := httpclient.Req(
// 		"https://"+host+resource+"?"+query,
// 	).AddHeader(
// 		"x-ms-date", date,
// 	).AddHeader(
// 		"x-ms-content-sha256", contentHash,
// 	).AddHeader(
// 		"Authorization", authorizationHeader,
// 	).AddHeader(
// 		"Content-Type", "application/json",
// 	).AddBody(
// 		reqbody,
// 	).Patch()
// 	responseBody := res.GetBody()
// 	if !res.IsSuccess() {
// 		return errors.New(string(responseBody))
// 	}
// 	if len(responseBody) == 0 {
// 		return nil
// 	}
// 	return json.Unmarshal(responseBody, &response)
// }
//
// func (c *client_) Posts(
// 	ctx context.Context,
// 	host string,
// 	resource string,
// 	query string,
// 	reqbody interface{},
// 	response interface{},
// ) error {
// 	body := []byte("{}")
// 	var err error
// 	if reqbody != nil {
// 		body, err = json.Marshal(reqbody)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		reqbody = struct{}{}
// 	}
// 	date := time.Now().UTC().Format(http.TimeFormat)
// 	// Compute a content hash for the 'x-ms-content-sha256' header.
// 	contentHash := computeContentHash(body)
// 	// Prepare a string to sign.
// 	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", resource+"?"+query, date, host, contentHash)
//
// 	// Compute the signature.
// 	signature := computeSignature(stringToSign, c.key)
//
// 	// Concatenate the string, which will be used in the authorization header.
// 	authorizationHeader := fmt.Sprintf(
// 		"HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s",
// 		signature,
// 	)
// 	res := httpclient.Req(
// 		"https://"+host+resource+"?"+query,
// 	).AddHeader(
// 		"x-ms-date", date,
// 	).AddHeader(
// 		"x-ms-content-sha256", contentHash,
// 	).AddHeader(
// 		"Authorization", authorizationHeader,
// 	).AddHeader(
// 		"Content-Type", "application/json",
// 	).AddBody(
// 		reqbody,
// 	).Post()
//
// 	responseBody := res.GetBody()
//
// 	fmt.Println("responseBody", string(responseBody))
//
// 	if !res.IsSuccess() {
// 		return errors.New(string(responseBody))
// 	}
//
// 	if len(responseBody) == 0 {
// 		return nil
// 	}
//
// 	return json.Unmarshal(responseBody, &response)
// }
//
// func (c *client_) Delete(
// 	ctx context.Context,
// 	host string,
// 	resource string,
// 	query string,
// 	response interface{},
// ) error {
// 	body := []byte("{}")
// 	reqBody := struct{}{}
// 	date := time.Now().UTC().Format(http.TimeFormat)
// 	contentHash, authHeader := createAuthHeader(
// 		"DELETE",
// 		host,
// 		resource+"?"+query,
// 		date,
// 		c.key,
// 		body,
// 	)
// 	res := httpclient.Req(
// 		"https://"+host+resource+"?"+query,
// 	).AddHeader(
// 		"x-ms-date", date,
// 	).AddHeader(
// 		"x-ms-content-sha256", contentHash,
// 	).AddHeader(
// 		"Authorization", authHeader,
// 	).AddHeader(
// 		"Content-Type", "application/json",
// 	).AddBody(reqBody).Del()
// 	responseBody := res.GetBody()
// 	if !res.IsSuccess() {
// 		return errors.New(string(responseBody))
// 	}
// 	if len(responseBody) == 0 {
// 		return nil
// 	}
// 	return json.Unmarshal(responseBody, &response)
// }
//
// func (c *client_) Get(
// 	ctx context.Context,
// 	host string,
// 	resource string,
// 	query string,
// 	response interface{},
// ) error {
// 	body := []byte("{}")
// 	reqBody := struct{}{}
// 	date := time.Now().UTC().Format(http.TimeFormat)
// 	contentHash, authHeader := createAuthHeader(
// 		"GET",
// 		host,
// 		resource+"?"+query,
// 		date,
// 		c.key,
// 		body,
// 	)
// 	res := httpclient.Req(
// 		"https://"+host+resource+"?"+query,
// 	).AddHeader(
// 		"x-ms-date", date,
// 	).AddHeader(
// 		"x-ms-content-sha256", contentHash,
// 	).AddHeader(
// 		"Authorization", authHeader,
// 	).AddHeader(
// 		"Content-Type", "application/json",
// 	).AddBody(reqBody).Get()
// 	responseBody := res.GetBody()
// 	if !res.IsSuccess() {
// 		return errors.New(string(responseBody))
// 	}
// 	if len(responseBody) == 0 {
// 		return nil
// 	}
// 	return json.Unmarshal(responseBody, &response)
// }
