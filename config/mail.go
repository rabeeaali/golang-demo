package config

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"strconv"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

type Data struct {
	Body string
}

//SendMail sends email
func SendMail(email string, code string, filename string, subject string) {
	// mail conditionals
	from := os.Getenv("MAIL_FROM_ADDRESS")
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	// set html files
	var t, _ = template.ParseFiles("email_templates/" + filename + ".html")
	var tpl bytes.Buffer
	data := Data{
		Body: code, // the data of email
	}
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}
	result := tpl.String()

	// set email config
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetAddressHeader("Cc", from, "ASU")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)
	d := gomail.NewDialer(host, port, username, password)

	// send email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
