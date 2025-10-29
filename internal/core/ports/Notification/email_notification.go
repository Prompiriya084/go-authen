package ports_notification

type EmailNotification interface {
	Notify(bodyMessage string) error
}
