package emails

type Payload struct {
	Headers                        Headers      `json:"headers"`
	SenderAddress                  string       `json:"senderAddress"`
	Content                        Content      `json:"content"`
	Recipients                     Recipients   `json:"recipients"`
	Attachments                    []Attachment `json:"attachments"`
	ReplyTo                        []ReplyTo    `json:"replyTo"`
	UserEngagementTrackingDisabled bool         `json:"userEngagementTrackingDisabled"`
}

type Attachment struct {
	Name            string `json:"name"`
	ContentType     string `json:"contentType"`
	ContentInBase64 string `json:"contentInBase64"`
}

type Content struct {
	Subject   string `json:"subject"`
	PlainText string `json:"plainText"`
	HTML      string `json:"html"`
}

type Headers struct {
	ClientCorrelationID    string `json:"ClientCorrelationId"`
	ClientCustomHeaderName string `json:"ClientCustomHeaderName"`
}

type Recipients struct {
	To  []ReplyTo `json:"to"`
	Cc  []ReplyTo `json:"cc"`
	Bcc []ReplyTo `json:"bcc"`
}

type ReplyTo struct {
	Address     string `json:"address"`
	DisplayName string `json:"displayName"`
}

type EmailResult struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

const (
	defaultAPIVersion = "2023-03-31"
)
