package main

import (
	"flag"
	"fmt"
	"net/smtp"
)

func main() {
	to := flag.String("to", "", "The email address to send to")
	message := flag.String("message", "", "The message to include in the email")
	flag.Parse()

	if *to == "" || *message == "" {
		fmt.Println("Please provide a 'to' and 'message' argument")
		return
	}

	from := ""
	pass := ""

	msg := fmt.Sprintf("<html><body>%s</body></html>", *message)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{*to}, []byte(msg))

	if err != nil {
		fmt.Printf("There was an error sending the email: %s\n", err)
		return
	}

	fmt.Println("Email sent successfully")
}
