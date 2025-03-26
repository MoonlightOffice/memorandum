package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func PubSub() {
	nc := InitConn()
	defer nc.Close()

	subject := "user.usr_asdf1234"

	go func() {
		count := 0
		for {
			err := nc.Publish(subject, fmt.Appendf([]byte{}, `{ "count": %d }`, count))
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Millisecond * 500)

			count++
		}
	}()

	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
		msg.AckSync()

		type Msg struct {
			Count int `json:"count"`
		}

		var m Msg
		json.Unmarshal(msg.Data, &m)
		if m.Count > 5 {
			msg.Sub.Unsubscribe()
			return
		}
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for sub.IsValid() {
		time.Sleep(time.Second)
	}
}
