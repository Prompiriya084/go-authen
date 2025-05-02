package ports_notification

import (
	entities "github.com/Prompiriya084/go-authen/Internal/Core/Entities"
)

type EmailNotification interface {
	Send(n *entities.NotificationMessage) error
}
