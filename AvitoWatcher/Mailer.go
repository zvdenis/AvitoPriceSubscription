package AvitoWatcher

import (
	"errors"
	gomail "gopkg.in/mail.v2"
	"regexp"
)

//Выражение проверки почты
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
const maxLen = 100
const minLen = 4

//Данные почты на сервисе MailGun
const from = "postmaster@sandbox21685e865fc44632a93dc8a063a11179.mailgun.org"
const password = "74a9713530b93b359ecd3aeeb1b94e07-cb3791c4-0857e7b1"
const smtpHost = "smtp.mailgun.org"
const smtpPort = 587

//Проверяет корректность почты
func IsEmailValid(e string) bool {
	if len(e) < minLen && len(e) > maxLen {
		return false
	}
	return emailRegex.MatchString(e)
}

//Отправляет почту списку адресатов, с заданным сообщением
func SendMail(to []string, text string) error {

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	for _, el := range to {
		m := gomail.NewMessage()
		m.SetHeader("From", from)
		m.SetHeader("To", el)
		m.SetHeader("Subject", "Avito watcher")
		m.SetBody("text/plain", text)

		if err := d.DialAndSend(m); err != nil {
			return errors.New("mail deliver failed")
		}
	}

	return nil
}