package smtp

import (
	"net/smtp"
	"exam/config"
)


func SendMail(toEmail string, msg string) error {

	from := "alishersamadov0625@gmail.com"
	to := []string{toEmail}
	subject := "Register for: EXAM "
	message := msg


	body := "To: " + to[0] + "\r\n" +
	"Subject: " + subject + "\r\n" +
	"\r\n" + message


	auth := smtp.PlainAuth("", config.SmtpUsername, config.SmtpPassword, config.SmtpServer)
 
	err := smtp.SendMail(config.SmtpServer+":"+config.SmtpPort, auth, from, to, []byte(body))
	if err != nil {
		return err
	}

	return nil
}