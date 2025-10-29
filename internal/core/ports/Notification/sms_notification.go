package ports_notification

type SMSNotification interface {
	Notify(bodyMessage string) error
}
