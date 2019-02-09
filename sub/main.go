package main

import (
	"fmt"
	"log"
	"runtime"

	nats "github.com/nats-io/go-nats"
)

func printMsg(m *nats.Msg) {
	log.Printf("Received on [%s]: '%s'\n", m.Subject, string(m.Data))
}

func main() {

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	nc.Subscribe("ch", func(msg *nats.Msg) {
		printMsg(msg)
		fmt.Println("Sub")
	})

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", "ch")

	runtime.Goexit()
}
