package main

import (
	"fmt"
	"os"
	"os/user"
)

func GetCurrentDirectory() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return pwd
}

func GetHostName() string {
	hname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hname
}

func GetUserName() string {
	uname, err := user.Current()
	if err != nil {
		return ""
	}
	return uname.Username
}

func CustomPrompt() string {
	username := GetUserName()
	hostname := GetHostName()
	wd := GetCurrentDirectory()

	return fmt.Sprintf("%s@%s:%s", username, hostname, wd)
}
