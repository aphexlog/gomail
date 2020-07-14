package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	gomail "gopkg.in/gomail.v2"
)

// Config type intended to authenticate to SMTP server
type Config struct {
	Username   string `json:"username"`
	Password   string `json:"password"` // use a properly hashed password (bcrypt / scrypt)
	Sender     string `json:"sender"`
	Recepients string `json:"recepients"`
	CC         string `json:"cc"`
}

func main() {

	// open json file
	jsonFile, err := os.Open("config.json")
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened auth.json")
	// defer the closing of the jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Config

	json.Unmarshal(byteValue, &config)
	fmt.Println("Attempting to send mail from " + config.Sender + "...")

	m := gomail.NewMessage()
	m.SetHeader("From", config.Sender)
	m.SetHeader("To", config.Recepients)
	m.SetAddressHeader("Cc", config.CC, "Aaron W. West")
	m.SetHeader("Subject", "Hello dumbdumb!")
	m.SetBody("text/html", "Hello <b>A-aron</b>")

	d := gomail.NewDialer("smtp.gmail.com", 587, config.Username, config.Password)

	fmt.Println("Username: " + config.Username)
	fmt.Println("Password: " + config.Password)
	fmt.Println("Sender: " + config.Sender)
	fmt.Println("Recepients: " + config.Recepients)
	fmt.Println("CC: " + config.CC)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
