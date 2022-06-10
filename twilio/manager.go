package twilio

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TwilioSMSManager struct {
	AccountSID  string
	AuthToken   string
	PhoneNumber string
	SendSmsURL  string
}

func NewTwilioConfig() *TwilioSMSManager {

	var SendSmsURL strings.Builder

	AccountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	AuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	PhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	SendSmsURL.WriteString("https://api.twilio.com/2010-04-01/Accounts/")
	SendSmsURL.WriteString(AccountSID)
	SendSmsURL.WriteString("/Messages.json")

	return &TwilioSMSManager{
		AccountSID:  AccountSID,
		AuthToken:   AuthToken,
		PhoneNumber: PhoneNumber,
		SendSmsURL:  SendSmsURL.String(),
	}
}

func (c *TwilioSMSManager) SendSmsRequest(to, body string) (*http.Request, error) {
	values := url.Values{}
	values.Set("To", to)
	values.Set("From", c.PhoneNumber)
	values.Set("Body", body)
	reader := strings.NewReader(values.Encode())
	req, err := http.NewRequest("POST", c.SendSmsURL, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.AccountSID, c.AuthToken)
	return req, nil
}

func (c *TwilioSMSManager) MustSendSmsRequest(to, body string) *http.Request {
	req, err := c.SendSmsRequest(to, body)
	if err != nil {
		log.Fatal(err)
	}
	return req
}
