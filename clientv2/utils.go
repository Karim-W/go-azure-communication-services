package clientv2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

const signed_header_prefix = "HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature="

func createAuthHeader(
	method string,
	host string,
	resourcePath string,
	date string,
	secret string,
	body []byte,
) (string, string) {
	contentHash := computeContentHash(body)
	stringToSign := stringBuilder(
		method,
		"\n",
		resourcePath,
		"\n",
		date,
		";",
		host,
		";",
		contentHash,
	)
	signature := computeSignature(stringToSign, secret)

	// Concatenate the string, which will be used in the authorization header.
	authorizationHeader := signed_header_prefix + signature

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

func stringBuilder(strs ...string) string {
	buff := strings.Builder{}
	for _, str := range strs {
		buff.WriteString(str)
	}
	return buff.String()
}
