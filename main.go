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
	from := readString()

	fmt.Print("Password: ")
	password := readString()

	to := []string{}

	i := 0
	for i != 1 {
		fmt.Print("Who do you want to send an email to? ")
		to = append(to, readString())
		fmt.Print("Anyone else (yes/no)? ")
		d := readString()
		if d != "yes" {
			i = 1
		}
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	fmt.Print("What type of message do you want to send: \n",
		"	1. Typed, \n",
		"	2. HTML template \n")
	typeOfMessage := readString()
	typeOfMessageI, err := strconv.Atoi(typeOfMessage)
	errorHandling(err)

	var body bytes.Buffer
	var message []byte

	if typeOfMessageI == 1 {
		fmt.Print("What message do you want to write? ")
		messageS := readString()
		message = []byte(messageS)
	} else if typeOfMessageI == 2 {
		t, _ := template.ParseFiles("template.html")
		fmt.Print("What's the subject of email? ")
		subject := readString()

		mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body.Write([]byte(fmt.Sprintf("Subject: "+subject+"\n%s\n\n", mimeHeaders)))

		t.Execute(&body, struct {
		}{})
		message = body.Bytes()
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	errorHandling(err)
	fmt.Println("Email Sent Successfully!")
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

func readString() string {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	errorHandling(err)
	b := strings.TrimSpace(a)
	return b
}
