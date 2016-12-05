// An example bot.

package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/super-project/slacker"
)

var token = "token"
var iamRegexp = regexp.MustCompile(`.*[Ii]'m\s+(\S+).*`)

func main() {
	s := slacker.NewSlacker(token)
	s.Connect()

	for {
		msg, _ := s.GetMessage()
		fmt.Println(msg)
		if strings.Contains(msg.Text, "'m") {
			matches := iamRegexp.FindAllStringSubmatch(msg.Text, 1)
			if len(matches) > 0 {
				name := matches[0][1]
				s.SendMessage("Hi "+name+", I'm DadJokeBot.", msg.Channel)
			}
		} else if strings.Contains(msg.Text, "ell me a joke") {
			joke := strings.Replace(Joke(), "?", ".", -1)
			lines := strings.Split(joke, ".")
			for _, line := range lines {
				s.SendMessage(line, msg.Channel)
				time.Sleep(2000 * time.Millisecond)
			}
		}
	}
}
