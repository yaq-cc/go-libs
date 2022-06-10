package twilio

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TwilioSMSClient struct {
	*http.Client
	AccountSID  string
	AuthToken   string
	PhoneNumber string
}

func NewTwilioSMSClient() *TwilioSMSClient {
	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	phoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	twilioSMSTransport := &twilioSMSTransport{
		AccountSID: accountSID,
		AuthToken:  authToken,
	}
	return &TwilioSMSClient{
		AccountSID:  accountSID,
		AuthToken:   authToken,
		PhoneNumber: phoneNumber,
		Client: &http.Client{
			Transport: twilioSMSTransport,
		},
	}
}

func (c *TwilioSMSClient) SendSMSURL() string {
	var sendUrl strings.Builder
	sendUrl.WriteString("https://api.twilio.com/2010-04-01/Accounts/")
	sendUrl.WriteString(c.AccountSID)
	sendUrl.WriteString("/Messages.json")
	return sendUrl.String()
}

func (c *TwilioSMSClient) NewSendSMSRequest(to, body string) *TwilioSendSMSRequest {
	return &TwilioSendSMSRequest{
		To:         to,
		From:       c.PhoneNumber,
		Body:       body,
		SendSMSURL: c.SendSMSURL(),
	}
}

func (c *TwilioSMSClient) SendSMS(r *TwilioSendSMSRequest) (*http.Response, error) {
	values := url.Values{}
	values.Set("To", r.To)
	values.Set("From", r.From)
	values.Set("Body", r.Body)
	reader := strings.NewReader(values.Encode())
	smsURL, err := url.Parse(r.SendSMSURL)
	if err != nil {
		return nil, err
	}
	queryVals := smsURL.Query()
	if r.Parameters != nil {
		for key, val := range r.Parameters {
			queryVals.Set(key, val)
		}
	}
	req, err := http.NewRequest("POST", r.SendSMSURL, reader)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type twilioSMSTransport struct {
	AccountSID string
	AuthToken  string
}

func (t *twilioSMSTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.AccountSID, t.AuthToken)
	return http.DefaultTransport.RoundTrip(r)
}

type TwilioSendSMSRequest struct {
	To         string
	From       string
	Body       string
	SendSMSURL string
	Parameters map[string]string
}

func (r *TwilioSendSMSRequest) Set(key, value string) {
	if r.Parameters == nil {
		r.Parameters = make(map[string]string)
	}
	r.Parameters[key] = value
}
