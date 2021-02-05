package main

import (
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
	"flag"
)

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
		m := gomail.NewMessage()
		m.SetHeader("From", "blockchainwarning@omnisolu.com")
		m.SetHeader("To", "folowal757@loopsnow.com", "blockchainwarning@omnisolu.com")
		//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", "Dappley Error Report:")
		m.SetBody("text/html", "<p>" + ERROR + "</p>")
		m.Attach(filename)

		d := gomail.NewDialer("smtp.gmail.com", 587, "blockchainwarning@omnisolu.com", "01353751")

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}
}