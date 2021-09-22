package notification

import (
	"errors"
	"fmt"

	"github.com/davidpolme/IntermediateGo/notification/types"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() types.ISender
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {

	switch notificationType {
	case "SMS":
		return &types.SMSNotification{}, nil
	case "Email":
		return &types.EmailNotification{}, nil
	default:
		return nil, errors.New("unsuported notification type")
	}
}

func SendNotification(factory INotificationFactory) {
	factory.SendNotification()
}

func GetMethod(factory INotificationFactory) {
	fmt.Println(factory.GetSender().GetSenderMethod())
}
