package utils

import (
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "inggarnugrahaputra@gmail.com"
	password := "ctlh attv bery ngcl" // password aplikasi gmail

	// gmail SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// format email 
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// autentikasi
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// kirim email
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}
