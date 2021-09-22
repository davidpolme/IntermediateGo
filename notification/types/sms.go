package types

import "fmt"

type SMSNotification struct {
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
