package git

import (
	"os/exec"
)

func GetUsername() (string, error){
	username, err := exec.Command("git", "config", "user.name").Output()
	return string(username), err
}

func GetEmail() (string, error){
	email, err := exec.Command("git", "config", "user.email").Output()
	return string(email), err
}