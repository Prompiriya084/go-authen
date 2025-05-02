package notification

import (
	"fmt"

	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
	ports_notification "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Notification"
)

type EmailNotificationImpl struct {
	Sender string
}

func NewEmailNotification(defaultSender string) ports_notification.EmailNotification {
	return &EmailNotificationImpl{Sender: defaultSender}
}

func (n *EmailNotificationImpl) Send(message *entities.NotificationMessage) error {
	fmt.Println(n.Sender)
	return nil
}
