package main

// An example implementation of a slacker bot

import (
	"fmt"
	"strings"
)

var token = "Token Goes Here"

func main() {
	s := NewSlacker(token)
	s.Connect()

	fmt.Println(s.ListChannels())
	for {
		msg, _ := s.GetMessage()
		fmt.Println(msg)
		if strings.Contains(msg.Text, "good boy") {
			fmt.Println(s.SendMessage("woof!!!", msg.Channel))
		}
	}
}
