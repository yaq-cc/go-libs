package twilio

import (
	"encoding/xml"
	"fmt"
	"io"
)

type Message string

type MessagingResponse struct {
	XMLName  xml.Name  `xml:"Response"`
	Say      string    `xml:"Say,omitempty"`
	Messages []Message `xml:"Message,omitempty"`
	Redirect string    `xml:"Redirect,omitempty"`
}

func NewMessagingResponse() *MessagingResponse {
	return &MessagingResponse{}
}

func (r *MessagingResponse) AddMessage(msg string) *MessagingResponse {
	r.Messages = append(r.Messages, Message(msg))
	return r
}

func (r *MessagingResponse) AddRedirect(url string) *MessagingResponse {
	r.Redirect = url
	return r
}

func (r *MessagingResponse) SendTo(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "\t")
	fmt.Fprint(w, xml.Header)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	return nil
}
