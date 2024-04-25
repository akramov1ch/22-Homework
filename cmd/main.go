package main

import (
	gt "User/git"
	"fmt"
	"log"
	"os"
)

func main() {
	username, err := gt.GetUsername()
	if err != nil {
		fmt.Println("Username ni o`qishda xatolik yuz berdi:")
		log.Fatal(err)
	}
	file, err := os.OpenFile("storage/username.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err!= nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(username)
	if err!= nil {
		log.Fatal(err)
	}
	email, err := gt.GetEmail()
	if err!= nil {
		fmt.Println("Email ni o`qishda xatolik yuz berdi:")
		log.Fatal(err)
	}
	_, err = file.WriteString("\n")
	if err!= nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(email)
	if err!= nil {
		log.Fatal(err)
	}
}
