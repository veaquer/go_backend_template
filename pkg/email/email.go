package email

type Email struct {
	To string
	Subject string
	Body string
	isHTML bool
}

type Sender interface {
	SendEmail(email Email) error
}
