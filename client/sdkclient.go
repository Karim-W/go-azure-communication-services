package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
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

func (c *Client) Post(
	earl string,
	body interface{},
	response interface{},
) error {
	urlObj, err := url.Parse(earl)
	if err != nil {
		return err
	}
	marshalledBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	date := time.Now().UTC().Format(time.RFC1123)
	host := urlObj.Host
	hash := c.computeSHA256Hash(marshalledBody)
	stringToSign := "POST\n" + earl + "\n" + date + ";" + host + ";" + string(hash)
	signature := c.hmacSHA256Hash([]byte(stringToSign))
	header := "HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=" + signature
	res := httpclient.Req(earl).
		AddHeader("x-ms-date", date).
		AddHeader("x-ms-content-sha256", hex.EncodeToString(hash)).
		AddHeader("Authorization", header).
		AddBody(body).
		Post()
	if !res.IsSuccess() {
		return res.CatchError()
	}
	err = res.SetResult(response)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) computeSHA256Hash(content []byte) []byte {
	hash := sha256.New()
	hash.Write(content)
	return hash.Sum(nil)
}

func (c *Client) hmacSHA256Hash(
	content []byte,
) string {
	hash := hmac.New(sha256.New, []byte(c.key))
	hash.Write(content)
	return hex.EncodeToString(hash.Sum(nil))
}
