package service

import (
	"fmt"
	"inine-track/pkg/config"
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", config.SMTP.From)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		config.SMTP.Host,
		config.SMTP.Port,
		config.SMTP.User,
		config.SMTP.Password,
	)

	if err := dialer.DialAndSend(msg); err != nil {
		log.Printf("Failed to send email to %s: %v", to, err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Printf("Email successfully sent to %s", to)
	return nil
}

func SendTokenEmail(email, token string) error {
	subject := "Seu Token de Acesso"
	body := fmt.Sprintf(`
	"<body style='font-family: Arial, sans-serif; font-size: 15px;'>"
		<h1>Seu Token de Acesso</h1>
		<p>Use o token abaixo para autenticar:</p>
		<code>%s</code>
		<p>Este token expirará em 1 hora.</p>
		<p>Se você não solicitou este token, por favor ignore este email.</p>
	`, token)

	return SendEmail(email, subject, body)
}
