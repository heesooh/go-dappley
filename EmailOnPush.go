package main

import (
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
	"bufio"
	"flag"
)

func main() {
	var fileName, commitInfo, commitEmail, senderEmail, senderPasswd string
	flag.StringVar(&fileName,     "fileName",     "log.txt",         "Test result file")
	flag.StringVar(&commitInfo,   "commitInfo",   "commit info",     "default commit info")
	flag.StringVar(&commitEmail,  "commitEmail",  "committer email", "default commit email")
	flag.StringVar(&senderEmail,  "senderEmail",  "sender_username@example.com", "Email address of the sender")
	flag.StringVar(&senderPasswd, "senderPasswd", "PASSWORD",        "Password of the sender's email address.")
	flag.Parse()

	sendEmail(fileName, commitInfo, commitEmail, senderEmail, senderPasswd)
}

func sendEmail(fileName string, commitInfo string, commitEmail string, senderEmail string, senderPasswd string) {
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
		errorEmail := "<p>"
		CommitScanner := bufio.NewScanner(strings.NewReader(COMMIT))
		for CommitScanner.Scan() {
			MSG := CommitScanner.Text()
			MSG = "<br>" + MSG
			errorEmail += MSG
		}
		errorEmail += "</p> <p>Failing Tests:"
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
		m.SetHeader("From", senderEmail)
		m.SetHeader("To", commitEmail)
		//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", "go-dappley Error Report:")
		m.SetBody("text/html", errorEmail)
		m.Attach(fileName)
		m.Attach("change.txt")

		d := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, senderPasswd)

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}