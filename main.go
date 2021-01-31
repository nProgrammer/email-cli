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

	fmt.Print("What message do you want to write? ")
	e, _ := r.ReadString('\n')
	messageS := strings.TrimSpace(e)
	message := []byte(messageS)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
