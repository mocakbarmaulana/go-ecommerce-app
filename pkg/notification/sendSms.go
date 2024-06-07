package notification

import (
	"encoding/json"
	"fmt"
	"github.com/go-ecommerce-app/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSms(phoneNumber string, message string) error
}

type notificationClient struct {
	c config.AppConfig
}

func NewNotificationClient(c config.AppConfig) NotificationClient {
	return &notificationClient{c: c}
}

func (n *notificationClient) SendSms(phoneNumber string, message string) error {
	// Send SMS using Twilio
	accountSid := n.c.TwilioAccountSid
	authToken := n.c.TwilioAuthToken
	phone := n.c.TwilioFromNumber

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	fmt.Println("phone", phone)

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(phone)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)

	if err != nil {
		fmt.Println("Error sending SMS: ", err.Error())
	} else {
		response, _ := json.Marshal(resp)
		fmt.Println("SMS sent successfully: ", string(response))
	}

	return err
}
