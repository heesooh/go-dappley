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
	//read log file
	errMSG, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//read commitInfo file
	commitMSG, err := ioutil.ReadFile(commitInfo)
	if err != nil {
		return
	}
	//convert to string
	ERROR := string(errMSG)
	COMMIT := string(commitMSG)

	if (strings.Contains(ERROR, "FAIL")) {
		//create the html formatted email
		errorEmail := "<p> Failing Tests:"
		CommitScanner := bufio.NewScanner(strings.NewReader(COMMIT))
		for CommitScanner.Scan() {
			MSG := CommitScanner.Text()
			MSG = "<br>" + MSG
			errorEmail += MSG
		}
		errorEmail += "</p> <p>"
		ErrorScanner := bufio.NewScanner(strings.NewReader(ERROR))
		for ErrorScanner.Scan() {
			MSG := ErrorScanner.Text()
			if (strings.Contains(MSG, "FAIL")) {
				MSG = "<br>" + MSG
				errorEmail += MSG
			}
		}
		errorEmail += "</p>"
		//send the email
		m := gomail.NewMessage()
		m.SetHeader("From", "blockchainwarning@omnisolu.com")
		m.SetHeader("To", "blockchainwarning@omnisolu.com")
		//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", "go-dappley Error Report:")
		m.SetBody("text/html", errorEmail)
		m.Attach(fileName)
		m.Attach("change.txt")

		d := gomail.NewDialer("smtp.gmail.com", 587, "blockchainwarning@omnisolu.com", "01353751")

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}