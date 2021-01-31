package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func main() {
	fmt.Print("Login: ")
	r := bufio.NewReader(os.Stdin)
	a, _ := r.ReadString('\n')
	from := strings.TrimSpace(a)

	fmt.Print("Password: ")
	b, _ := r.ReadString('\n')
	password := strings.TrimSpace(b)
	// Sender data.

	// Receiver email address.
	to := []string{
		"wagnernorbert836@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
