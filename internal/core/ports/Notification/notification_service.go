package ports_notification

type NotificationService interface {
	Notify(body string) error
}
