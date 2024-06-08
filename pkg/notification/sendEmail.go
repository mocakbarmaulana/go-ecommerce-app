package notification

import (
	"github.com/go-ecommerce-app/config"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

type EmailClient interface {
	SendEmail(email string, subject string, message string) error
}

type emailClient struct {
	cfg config.AppConfig
}

func NewEmailClient(c config.AppConfig) EmailClient {
	return &emailClient{cfg: c}
}

func (e *emailClient) SendEmail(email string, subject string, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.cfg.MailFrom)
	m.SetHeader("To", email)
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	port, _ := strconv.Atoi(e.cfg.MailPort)

	d := gomail.NewDialer(e.cfg.MailHost, port, e.cfg.MailUsername, e.cfg.MailPassword)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Could not send email: %v", err)
		return err
	}

	return nil
}
