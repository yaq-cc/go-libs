package twilio

import (
	"net/http"
	"os"
	"testing"
)

func TestTwilioConfig(t *testing.T) {
	cfg := NewTwilioConfig()
	t.Log(cfg.AccountSID)
	t.Log(cfg.AuthToken)
	t.Log(cfg.PhoneNumber)
	t.Log(cfg.SendSmsURL)
}

func TestTwilioSendSms(t *testing.T) {
	phoneNumber := os.Getenv("TEST_PHONE_NUMBER")
	cfg := NewTwilioConfig()
	req := cfg.MustSendSmsRequest(phoneNumber, "Yvan says hi.")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.StatusCode)
}
