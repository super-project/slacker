
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/super-project/slacker"
)

var token = "TOKEN_GOES_HERE"
var iamRegexp = regexp.MustCompile(`.*[Ii]'m\s+(\S+).*`)

func main() {
	s := slacker.NewSlacker(token)
	s.Connect()

	for {
		msg, _ := s.GetMessage()
		fmt.Println(msg)
		if strings.Contains(msg.Text, "xkcd") {
			xkcd := strings.Replace(Xkcd(), "\\", "", -1)
			// lines := strings.Split(xkcd, ".")
			// for _, line := range lines {
			// 	s.SendMessage(line, msg.Channel)
			// }
			s.SendMessage(xkcd, msg.Channel)
		}
	}
}
