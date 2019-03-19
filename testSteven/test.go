package main

import (
	"log"
)


func main() {
	Cstring := "Get me the weather"

	var commands = []string {
		"Get me the weather",
		"Events near me",
		"Send Email",
	}
	//if string is equal to command
	i := 0
	for i < 3 {
		if Compare(Cstring, commands[i]) == 0 {
			log.Println(commands[i])
		}
	}
}