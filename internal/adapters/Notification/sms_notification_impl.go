package notification

import (
	"fmt"

	ports_notification "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Notification"
)

type smsNotificationServiceImpl struct {
	Sender   string
	Receiver []*string
}

func NewSMSNotificationService(defaultSender string, defaultReceiver []*string) ports_notification.SMSNotification {
	return &smsNotificationServiceImpl{
		Sender:   defaultSender,
		Receiver: defaultReceiver,
	}
}
func (s *smsNotificationServiceImpl) Notify(body string) error {
	fmt.Println("Send sms notification to: %s from %s", s.Receiver, s.Sender)
	return nil
}
