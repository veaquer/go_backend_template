package email

import (
	"backend_template/internal/config"
	"fmt"

	"gopkg.in/mail.v2"
)

type GoMailSender struct {
	from string
	dialer *mail.Dialer
}

func NewGoMailSender(cfg config.EmailConfig) *GoMailSender {
	dialer := mail.NewDialer(cfg.Host,cfg.Port,cfg.Username,cfg.Password)
	return &GoMailSender{
		from:cfg.From,
		dialer:dialer,
	}
}

func (s *GoMailSender) SendEmail(e Email) error {
	m := mail.NewMessage()
	m.SetHeader("From",s.from)
	m.SetHeader("To",e.To)
	m.SetHeader("Subject",e.Subject)

	if e.isHTML {
		m.SetBody("text/html",e.Body)
	} else {
		m.SetBody("text/plain",e.Body)
	}

	if err := s.dialer.DialAndSend(m);err != nil{
		return fmt.Errorf("failed to send email: %w",err)
	}

	return nil
}
