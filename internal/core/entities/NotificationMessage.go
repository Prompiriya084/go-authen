package entities

type NotificationMessage struct {
	from string
	to   string
	body string
}

type EmailMessage struct {
	NotificationMessage
	cc      []string
	subject string
	footer  string
}
