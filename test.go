package main

import (
	"log"
	"strings"
)


func main() {
	Cstring := "Get me the weather"

	var commands = []string {
		"Get me the weather",
		"Events near me",
		"Send Email",
	}
	//if string is equal to command
	for i := 0; i < 3; i++ {
		if strings.Compare(Cstring, commands[i]) == 0 {
			log.Println(commands[i])
		}
	}
}