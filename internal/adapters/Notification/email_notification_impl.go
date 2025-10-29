package notification

import (
	"fmt"

	ports_notification "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Notification"
)

type emailNotificationImpl struct {
	Sender   string
	Receiver []*string
	Cc       []*string
}

func NewEmailNotification(defaultSender string, defaultReceiver []*string, defaultCC []*string) ports_notification.EmailNotification {
	return &emailNotificationImpl{
		Sender:   defaultSender,
		Receiver: defaultReceiver,
		Cc:       defaultCC,
	}
}

func (n *emailNotificationImpl) Notify(bodyMessage string) error {
	fmt.Println(n.Sender)
	return nil
}
