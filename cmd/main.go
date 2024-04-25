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
	file, err := os.OpenFile("username.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err!= nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(username)
	if err!= nil {
		log.Fatal(err)
	}
}
