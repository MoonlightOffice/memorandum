package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	KB = 1024
	MB = 1024 * KB
)

func SetupQueue() (streamInfo *nats.StreamInfo, consumerInfo *nats.ConsumerInfo, cleanup func(), err error) {
	js, close := InitJetStream()
	defer close()

	streamInfo, err = js.AddStream(&nats.StreamConfig{
		Name:       "sample-stream",
		Subjects:   []string{"users.usr_1", "users.usr_2", "companies.comp_1"},
		Storage:    nats.FileStorage,
		Replicas:   3,
		Retention:  nats.WorkQueuePolicy,
		Discard:    nats.DiscardNew,
		MaxMsgSize: 10 * MB,
	})
	if err != nil {
		return nil, nil, nil, err
	}

	consumerInfo, err = js.AddConsumer(streamInfo.Config.Name, &nats.ConsumerConfig{
		Name:      "sample-consumer",
		AckPolicy: nats.AckExplicitPolicy,
		AckWait:   5 * time.Minute,
		// Replicas:  3, // Inherits stream's replicas by default
		// MaxDeliver: 3, // DLQ -> $JS.EVENT.ADVISORY.CONSUMER.MAX_DELIVERIES.<STREAM>.<CONSUMER>
	})
	if err != nil {
		return nil, nil, nil, err
	}

	return streamInfo, consumerInfo, func() {
		js, close := InitJetStream()
		defer close()

		if err := js.DeleteStream(streamInfo.Config.Name); err != nil {
			fmt.Println(err)
		}
	}, nil
}

func QueuePub() {
	streamInfo, _, cleanup, err := SetupQueue()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cleanup()

	subject := streamInfo.Config.Subjects[0]
	fmt.Println("Publishing to", subject)

	js, close := InitJetStream()
	defer close()

	count := 0
	for {
		msg := fmt.Appendf([]byte{}, `{ "count": %d }`, count)
		_, err := js.Publish(subject, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(msg))
		time.Sleep(time.Second)

		count++
	}
}

func QueueSub() {
	streamInfo, consumerInfo, cleanup, err := SetupQueue()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cleanup()

	subject := streamInfo.Config.Subjects[0]
	fmt.Println("Subscribing", subject)

	js, close := InitJetStream()
	defer close()

	sub, err := js.PullSubscribe(subject, "", nats.Bind(streamInfo.Config.Name, consumerInfo.Name))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sub.Unsubscribe()

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	for {
		msgs, err := sub.Fetch(10, nats.Context(ctx))
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, msg := range msgs {
			err := msg.InProgress()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Message:", string(msg.Data))
			err = msg.AckSync()
			if err != nil {
				fmt.Println(err)
				return
			}

			//msg.Nak()
			//msg.Term()
		}

		time.Sleep(time.Second)
	}
}
