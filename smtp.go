package main

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	smtpServer = "smtp.***"
	smtpPort   = "587"
	email      = "***"
	password   = "***" // Лучше использовать переменные окружения
)

func SendEmail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", email, password, smtpServer)

	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, email, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	to := "***"
	subject := "Hello from Golang and me :)"
	body := "This is the test email sent from me to me"

	err := SendEmail(to, subject, body)

	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}
	fmt.Println("Email sent successfully")
}
