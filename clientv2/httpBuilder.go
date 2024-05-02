package clientv2

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/karim-w/stdlib/httpclient"
)

func (c *client_) buildRequest(
	method string,
	host string,
	resource string,
	queryParsms map[string][]string,
	reqbody any,
) (httpclient.HTTPRequest, error) {
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

	content, authHeader := createAuthHeader(
		method,
		host,
		resour,
		date,
		c.key,
		body,
	)

	res := httpclient.Req(
		"https://"+host+resour,
	).AddHeader(
		"x-ms-date", date,
	).AddHeader(
		"x-ms-content-sha256", content,
	).AddHeader(
		"Authorization", authHeader,
	).AddHeader(
		"Content-Type", "application/json",
	).AddBodyRaw(
		body,
	)

	return res, nil
}
