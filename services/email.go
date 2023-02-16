package services

import (
	"fmt"
	"influx-alert-api/global"
	"influx-alert-api/models"
	"log"
	"net/smtp"
	"strings"
)

func SendEmail(subject, message string) models.Response {

	res := models.Response{}
	res.Success = false

	env_smtp := global.EnvConfig.Email.Smtp
	env_user := global.EnvConfig.Email.User
	env_passwd := global.EnvConfig.Email.Password
	env_host := global.EnvConfig.Email.Host
	env_sender := global.EnvConfig.Email.Sender
	env_to := global.EnvConfig.Email.To

	var msgString string
	msgString += fmt.Sprintf("Subject: %s\r\n", subject)
	msgString += fmt.Sprintf("From: %s\r\n", global.EnvConfig.Email.Sender)
	msgString += fmt.Sprintf("To: %s\r\n", strings.Join(env_to, ";"))
	msgString += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	message = fmt.Sprintf("<p>%s</p>", message)
	message = strings.ReplaceAll(message, "\n", "</p><p>")
	msgString += message

	fmt.Println(msgString)

	msg := []byte(msgString)

	if global.EnvConfig.Email.Auth {
		auth := smtp.PlainAuth(env_smtp, env_user, env_passwd, env_host)

		err := smtp.SendMail(
			env_smtp,
			auth,
			env_sender,
			env_to,
			msg,
		)

		if err != nil {
			log.Printf("[Error] Failed to send email (Auth), err: %s\n", err)
			res.Msg = err.Error()
			return res
		}
	} else {
		err := smtp.SendMail(
			env_smtp,
			nil,
			env_sender,
			env_to,
			msg,
		)
		if err != nil {
			log.Printf("[Error] Failed to send email (No Auth), err: %s\n", err)
			// res.Msg = err.Error()
			// return res
		}
		c, err := smtp.Dial(env_smtp)
		if err != nil {
			log.Printf("[Error] Failed to send email (No Auth) by Dail, err:%s\n", err)
			res.Msg = err.Error()
			return res
		}
		defer c.Quit()

		// Set the sender and recipient.

		err = c.Mail(env_sender)

		if err != nil {
			log.Printf("[Error] Failed to send email (No Auth) by Mail, err:%s\n", err)
			res.Msg = err.Error()
			return res
		}

		for _, receiver := range env_to {
			err = c.Rcpt(receiver)
			if err != nil {
				log.Printf("[Error] Failed to send email (No Auth) by Rcpt, err:%s\n", err)
				res.Msg = err.Error()
				return res
			}
		}

		// Send the email body.
		wc, _ := c.Data()
		if err != nil {
			log.Printf("[Error] Failed to send email (No Auth) by Data, err:%s\n", err)
			res.Msg = err.Error()
			return res
		}

		defer wc.Close()

		_, err = wc.Write(msg)
		if err != nil {
			log.Printf("[Error] Failed to send email (No Auth) by WriteTo, err:%s\n", err)
			res.Msg = err.Error()
			return res
		}

	}

	res.Success = true
	res.Msg = "Success to send email"

	return res

}
