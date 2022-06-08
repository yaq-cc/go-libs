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
	cfg.Parameters["Owner"] = "Yvan"
	cfg.Parameters["Wife"] = "Wendy"
	req := cfg.MustSendSmsRequest(phoneNumber, "Yvan says hi.")
	t.Log(req.URL.String())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.StatusCode)
}
