package main

import (
	"fmt"
	"github.com/jaaaaaaaaaam/ircbot/modules/tvmaze"
	"github.com/joho/godotenv"
	"github.com/thoj/go-ircevent"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	roomName := os.Getenv("CHANS")

	con := irc.IRC(os.Getenv("NICK"), os.Getenv("USERNAME"))
	err = con.Connect(os.Getenv("NETWORK"))
	if err != nil {
		fmt.Println("Error connecting")
		return
	}

	con.AddCallback("001", func(e *irc.Event) {
		con.Join(roomName)
	})

	con.AddCallback("JOIN", func(e *irc.Event) {
		if e.Nick != os.Getenv("NICK") {
			//welcomeMsg := fmt.Sprintf("Welcome %s...", e.Nick)
			//con.Privmsg(roomName, welcomeMsg)
		}
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		// split is the message split into 2 parts
		// split[0] is the function that is matched
		// split[1] is the argument that is passed to the function.
		// For example !show girls
		split := strings.SplitAfterN(e.Message(), " ", 2)
		if strings.HasPrefix(split[0], "!") {
			command := strings.TrimPrefix(strings.TrimSpace(split[0]), "!")
			switch command {
			case "help":
				con.Privmsg(roomName, "help pls")
			case "show":
				msg := fmt.Sprintf("Looking up '%s'", split[1])
				con.Privmsg(e.Arguments[0], msg)
				ret := tvmaze.ShowLookup(split[1])
				con.Privmsg(e.Arguments[0], ret)
			default:
				con.Privmsg(e.Arguments[0], "NOPE")
			}
		}
	})

	con.Loop()
}
