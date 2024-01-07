package email

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
)

type EmailService struct {
	SmtpHost      string
	SmtpPort      string
	SmtpUser      string
	SmtpPass      string
	EmailTemplate *template.Template
	//Logger        utils.Logger
}

type Message struct {
	Header       string
	Username     string
	Introduction string
	Content      string
	URL          string
	ActionTitle  string
}

type Email struct {
	From      string
	Recipient string
	Subject   string
	Body      string
}

func NewEmailService() *EmailService {
	tmpl, err := template.ParseFiles("../static/index.html")
	if err != nil {
		//logger.Error("Failed to load email template", logrus.Fields{"function": "NewEmailService", "error": err.Error()})
	}

	return &EmailService{
		SmtpHost:      os.Getenv("SMTP_HOST"),
		SmtpPort:      os.Getenv("SMTP_PORT"),
		SmtpUser:      os.Getenv("SMTP_USER"),
		SmtpPass:      os.Getenv("SMTP_PASS"),
		EmailTemplate: tmpl,
		//Logger:        logger,
	}
}

func (e *EmailService) LoadEmail(message Message) (bytes.Buffer, error) {
	var body bytes.Buffer
	err := e.EmailTemplate.Execute(&body, message)
	if err != nil {
		//e.Logger.Error("Error executing template", logrus.Fields{"function": "LoadEmail", "error": err.Error()})
		return bytes.Buffer{}, err
	}
	return body, nil
}

func (e *EmailService) SendEmail(recipient, subject string, body bytes.Buffer) error {
	smtpAuth := smtp.PlainAuth("", e.SmtpUser, e.SmtpPass, e.SmtpHost)

	msg := []byte("Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body.String())

	err := smtp.SendMail(e.SmtpHost+":"+e.SmtpPort, smtpAuth, e.SmtpUser, []string{recipient}, msg)
	if err != nil {
		//e.Logger.Error("SMTP error while sending email", logrus.Fields{"function": "SendMail", "error": err.Error()})
		return err
	}

	return nil
}
