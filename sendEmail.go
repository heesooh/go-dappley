package main

import (
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
	"bufio"
	"flag"
)

func main() {
	var fileName string
	var commitInfo string
	var commitEmail string
	flag.StringVar(&fileName, "fileName", "default.txt", "default txt file")
	flag.StringVar(&commitInfo, "commitInfo", "default info", "default commit info")
	flag.StringVar(&commitEmail, "commitEmail", "default email", "default commit email")
	flag.Parse()
	sendEmail(fileName, commitInfo, commitEmail)
}

func sendEmail(fileName string, commitInfo string, commitEmail string) {
	errMSG, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	ERROR := string(errMSG)
	if (strings.Contains(ERROR, "FAIL")) {
		errorEmail := commitInfo + commitEmail + "<p> Failing Tests: </p>"
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
		m.Attach(fileName)

		d := gomail.NewDialer("smtp.gmail.com", 587, "blockchainwarning@omnisolu.com", "01353751")

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}