package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nuveo/speech-to-text"
)

var (
	inputFile = flag.String("input", "", "File to convert into text")
)

func main() {
	flag.Parse()

	if *inputFile != "" {
		c := speech.Credentials{}
		c.Setup()

		url := c.MakeSessionURL()
		sess, err := speech.GetSession(url)
		if err != nil {
			return
		}

		status, err := sess.GetRecognize()
		if err != nil {
			return
		}

		if status.Session.State != "initialized" {
			log.Println("Not ready yet, got:", status.Session.State)
			return
		}

		text, err := sess.SendAudio(*inputFile)
		if err != nil {
			return
		}

		fmt.Println(text)
		sess.DeleteSession()
	}
}
