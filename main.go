package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func main() {
	fmt.Print("Login: ")
	r := bufio.NewReader(os.Stdin)
	a, _ := r.ReadString('\n')
	from := strings.TrimSpace(a)

	fmt.Print("Password: ")
	b, _ := r.ReadString('\n')
	password := strings.TrimSpace(b)

	to := []string{}

	i := 0
	for i != 1 {
		fmt.Print("Who do you want to send an email to? ")
		c, _ := r.ReadString('\n')
		c = strings.TrimSpace(c)
		to = append(to, c)
		fmt.Print("Anyone else (yes/no)? ")
		d, _ := r.ReadString('\n')
		d = strings.TrimSpace(d)
		if d != "yes" {
			i = 1
		}
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	fmt.Print("What type of message do you want to send: \n",
		"	1. Typed, \n",
		"	2. HTML template \n")
	f, _ := r.ReadString('\n')
	typeOfMessage := strings.TrimSpace(f)
	typeOfMessageI, _ := strconv.Atoi(typeOfMessage)

	var body bytes.Buffer
	var message []byte

	if typeOfMessageI == 1 {
		fmt.Print("What message do you want to write? ")
		e, _ := r.ReadString('\n')
		messageS := strings.TrimSpace(e)
		message = []byte(messageS)
	} else if typeOfMessageI == 2 {
		t, _ := template.ParseFiles("template.html")
		fmt.Print("What's the subject of email? ")
		subject, _ := r.ReadString('\n')
		subject = strings.TrimSpace(subject)

		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: "+subject+"\n%s\n\n", mimeHeaders)))

		t.Execute(&body, struct {
		}{})
		message = body.Bytes()
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
