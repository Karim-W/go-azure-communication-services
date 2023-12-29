package sms

import (
	"context"
	"os"
	"testing"
)

func TestSMSClient(t *testing.T) {
	host := os.Getenv("ACS_HOST")
	key := os.Getenv("ACS_KEY")
	recipient := os.Getenv("SMS_RECIPIENT")
	sender := os.Getenv("SMS_SENDER")
	if host == "" || key == "" || recipient == "" || sender == "" {
		t.Skip("Skipping test because environment variables are not set")
	}
	client := NewClient(host, key, nil)
	payload := Request{
		From:    sender,
		Message: "Hello from Go!",
		SMSRecipients: []SMSRecipient{
			{
				To: recipient,
			},
		},
	}

	_, err := client.SendSMS(context.TODO(), nil, payload)
	if err != nil {
		t.Fatal(err)
	}
}
