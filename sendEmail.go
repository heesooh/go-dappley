package main

import (
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
	"bufio"
	"flag"
)

	// commit hash number
	// commit date 
	// commit message
	// name of the person who commited

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "default.txt", "default txt file")
	flag.Parse()
	sendEmail(filename)
}

func sendEmail(filename string) {
	errMSG, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	ERROR := string(errMSG)
	if (strings.Contains(ERROR, "FAIL")) {
		errorEmail := "<p> Failing Tests: </p>"
		scanner := bufio.NewScanner(strings.NewReader(ERROR))
		for scanner.Scan() {
			MSG := scanner.Text()
			if (strings.Contains(MSG, "FAIL")) {
				MSG = "<p>" + MSG + "</p>"
				errorEmail = errorEmail + MSG
			}
		}
		m := gomail.NewMessage()
		m.SetHeader("From", "blockchainwarning@omnisolu.com")
		m.SetHeader("To", "blockchainwarning@omnisolu.com", "blockchainwarning@omnisolu.com")
		//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", "Dappley Error Report:")
		m.SetBody("text/html", errorEmail)
		m.Attach(filename)

		d := gomail.NewDialer("smtp.gmail.com", 587, "blockchainwarning@omnisolu.com", "01353751")

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}