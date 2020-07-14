package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	gomail "gopkg.in/gomail.v2"
)

// User type intended to authenticate to SMTP server
type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"` // use a properly hashed password (bcrypt / scrypt)
	Sender     string `json:"sender"`
	Recepients string `json:"recepients"`
	CC         string `json:"cc"`
}

func main() {

	// open json file
	jsonFile, err := os.Open("auth.json")
	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened auth.json")
	// defer the closing of the jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var user User

	json.Unmarshal(byteValue, &user)
	fmt.Println("Attempting to send mail from " + user.Sender + "...")

	m := gomail.NewMessage()
	m.SetHeader("From", user.Sender)
	m.SetHeader("To", user.Recepients)
	m.SetAddressHeader("Cc", user.CC, "Aaron W. West")
	m.SetHeader("Subject", "Hello dumbdumb!")
	m.SetBody("text/html", "Hello <b>A-aron</b>")

	d := gomail.NewDialer("smtp.gmail.com", 587, user.Username, user.Password)

	fmt.Println("Username: " + user.Username)
	fmt.Println("Password: " + user.Password)
	fmt.Println("Sender: " + user.Sender)
	fmt.Println("Recepients: " + user.Recepients)
	fmt.Println("CC: " + user.CC)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
