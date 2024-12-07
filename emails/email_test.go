package emails_test

import (
	"context"
	"os"
	"testing"

	"github.com/karim-w/go-azure-communication-services/emails"
)

func TestEmailClient(t *testing.T) {
	host := os.Getenv("ACS_HOST")
	key := os.Getenv("ACS_KEY")
	recipient := os.Getenv("EMAIL_RECIPIENT")
	sender := os.Getenv("EMAIL_SENDER")
	if host == "" || key == "" || recipient == "" || sender == "" {
		t.Skip("Skipping test because environment variables are not set")
	}
	client := emails.NewClient(host, key, nil)
	payload := emails.Payload{
		Headers: emails.Headers{
			ClientCorrelationID:    "1234",
			ClientCustomHeaderName: "ClientCustomHeaderValue",
		},
		SenderAddress: sender,
		Content: emails.Content{
			Subject:   "Test email",
			PlainText: "This is a test email",
		},
		Recipients: emails.Recipients{
			To: []emails.ReplyTo{
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
