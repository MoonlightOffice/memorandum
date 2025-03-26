package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func InitConn() *nats.Conn {
	nc, err := nats.Connect(
		"nats://localhost:4222",
		nats.Name("example.go-client"),
		nats.Timeout(5*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	return nc
}

func InitJetStream() (js nats.JetStreamContext, close func()) {
	nc := InitConn()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	return js, nc.Close
}
