package ports_notification

import entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"

type Notification interface {
	Send(message *entities.NotificationMessage) error
}
