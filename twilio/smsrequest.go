package twilio

import (
	"io"
	"net/http"
	"time"

	"github.com/ajg/form"
)

type MessagingRequest struct {
	ToCountry        string `form:"ToCountry,omitempty"`
	ToState          string `form:"ToState,omitempty"`
	SmsMessageSid    string `form:"SmsMessageSid,omitempty"`
	NumMedia         int    `form:"NumMedia,omitempty"`
	ToCity           string `form:"ToCity,omitempty"`
	FromZip          string `form:"FromZip,omitempty"`
	SmsSid           string `form:"SmsSid,omitempty"`
	FromState        string `form:"FromState,omitempty"`
	SmsStatus        string `form:"SmsStatus,omitempty"`
	FromCity         string `form:"FromCity,omitempty"`
	Body             string `form:"Body,omitempty"`
	FromCountry      string `form:"FromCountry,omitempty"`
	To               string `form:"To,omitempty"`
	ToZip            string `form:"ToZip,omitempty"`
	NumSegments      int    `form:"NumSegments,omitempty"`
	ReferralNumMedia int    `form:"ReferralNumMedia,omitempty"`
	MessageSid       string `form:"MessageSid,omitempty"`
	AccountSid       string `form:"AccountSid,omitempty"`
	ApiVersion       string `form:"ApiVersion,omitempty"`
	From             string `form:"From,omitempty"`
	ReceivedTime     int64
}

func (r *MessagingRequest) FromHttpRequest(req *http.Request) error {
	enc := form.NewDecoder(req.Body)
	err := enc.Decode(r)
	if err != nil {
		return err
	}
	r.ReceivedTime = time.Now().Unix()
	return nil
}

func (mr *MessagingRequest) FromReader(r io.Reader) error {
	enc := form.NewDecoder(r)
	err := enc.Decode(mr)
	if err != nil {
		return err
	}
	mr.ReceivedTime = time.Now().Unix()
	return nil
}
