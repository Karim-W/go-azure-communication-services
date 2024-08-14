package emails

import (
	"context"
	"os"
	"testing"
)

func TestEmailClient(t *testing.T) {
	host := os.Getenv("ACS_HOST")
	key := os.Getenv("ACS_KEY")
	recipient := os.Getenv("EMAIL_RECIPIENT")
	sender := os.Getenv("EMAIL_SENDER")
	if host == "" || key == "" || recipient == "" || sender == "" {
		t.Skip("Skipping test because environment variables are not set")
	}
	client := NewClient(host, key, nil)
	payload := Payload{
		Headers: Headers{
			ClientCorrelationID:    "1234",
			ClientCustomHeaderName: "ClientCustomHeaderValue",
		},
		SenderAddress: sender,
		Content: Content{
			Subject:   "Test email",
			PlainText: "This is a test email",
		},
		Recipients: Recipients{
			To: []ReplyTo{
				{
					Address: recipient,
				},
			},
		},
		Attachments: []Attachment{},
	}

	result, err := client.SendEmail(context.TODO(), payload)
	if err != nil {
		t.Fatal(err)
	}
	if result.ID == "" {
		t.Fatal("TrackingId is empty")
	}
	if result.Status == "" {
		t.Fatal("Status is empty")
	}
}
