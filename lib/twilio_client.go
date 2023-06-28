package lib

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)


type TwilioClientInterface interface {
	SendSMS(toPhoneNumber string, fromPhoneNumber string, message string) error
}

type TwilioClient struct {}

func (twilioClient *TwilioClient) SendSMS(toPhoneNumber string, fromPhoneNumber string, message string) error {
	client := twilio.NewRestClient()
	params := &twilioApi.CreateMessageParams{}

	params.SetBody(message)
	params.SetFrom(fromPhoneNumber)
	params.SetTo(toPhoneNumber)

	_, error := client.Api.CreateMessage(params)

	return error
}
