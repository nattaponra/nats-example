package main

import (
	"log"

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
	//defer nc.Close()
	msg := []byte("Hello")
	nc.Publish("ch", msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", "ch", msg)
	}

	nc.Flush()
}
