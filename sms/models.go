package sms

type Request struct {
	From           string         `json:"from"`
	SMSRecipients  []SMSRecipient `json:"smsRecipients"`
	Message        string         `json:"message"`
	SMSSendOptions SMSSendOptions `json:"smsSendOptions"`
}

type SMSRecipient struct {
	To                     string  `json:"to"`
	RepeatabilityRequestID *string `json:"repeatabilityRequestId,omitempty"`
	RepeatabilityFirstSent *string `json:"repeatabilityFirstSent,omitempty"`
}

type SMSSendOptions struct {
	EnableDeliveryReport bool   `json:"enableDeliveryReport"`
	Tag                  string `json:"tag"`
}

type Result struct {
	Value []SMSResult `json:"value"`
}

type SMSResult struct {
	To                  string  `json:"to"`
	MessageID           *string `json:"messageId,omitempty"`
	HTTPStatusCode      int64   `json:"httpStatusCode"`
	Successful          bool    `json:"successful"`
	ErrorMessage        *string `json:"errorMessage,omitempty"`
	RepeatabilityResult *string `json:"repeatabilityResult,omitempty"`
}

var defaultAPIVersion = "2021-03-07"
