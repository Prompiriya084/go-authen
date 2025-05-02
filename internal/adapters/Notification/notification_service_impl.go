package notification

import ports_notification "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Notification"

type notificationServiceImpl struct {
	notifications []ports_notification.Notification
}

func NewNotificationService(notifications []ports_notification.Notification) ports_notification.NotificationService {
	return &notificationServiceImpl{
		notifications: notifications,
	}
}
func (s *notificationServiceImpl) Notify(body string) error {
	for _, notification := range s.notifications {
		if err := notification.Send(body); err != nil {
			return err
		}
	}
	return nil
}
