package main

import (
	nt "github.com/davidpolme/IntermediateGo/notification"
)

func main() {
	//SMS || Email
	notificationFactory, err := nt.GetNotificationFactory("Email")
	if err != nil {
		panic(err)
	}
	nt.SendNotification(notificationFactory)
	nt.GetMethod(notificationFactory)
}
