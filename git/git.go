package git

import (
	"fmt"
	"log"
	"os/exec"
)

func GetUsername() (string, error){
	username, err := exec.Command("git", "config", "user.name").Output()
	return string(username), err
}