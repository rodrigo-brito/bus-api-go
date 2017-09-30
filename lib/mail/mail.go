package mail

import (
	"net/smtp"

	"fmt"

	"github.com/spf13/viper"
)

func SendMessage(from, subject, name, message string) error {
	settings := viper.GetStringMapString("mail")
	auth := smtp.PlainAuth(
		"",
		settings["username"],
		settings["password"],
		settings["host"],
	)

	header := fmt.Sprintf("To: %q \r\nSubject: %q \r\n\r\nName: %q (%s)\r\n",
		settings["username"], subject, name, from)

	return smtp.SendMail(
		fmt.Sprintf("%s:%s", settings["host"], settings["port"]),
		auth,
		settings["username"],
		[]string{settings["destiny"]},
		[]byte(header+message),
	)
}
