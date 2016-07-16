package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats"
)

func main() {

	host := flag.String("host", nats.DefaultURL, "Host Address")
	runType := flag.String("type", "", "Type of command ( publish, subscribe )")
	topic := flag.String("topic", "", "Topic to publish / subscribe")
	data := flag.String("data", "", "Data to publish")
	queuGroup := flag.String("queueGroup", "", "Queue group for subscription")
	flag.Parse()

	nc, err := nats.Connect(*host)
	if err != nil {
		log.Fatal("Failed connecting to nats server")

	} else {
		log.Print("Connected to nats at url: " + *host)
	}

	if *runType == "publish" {
		if *topic == "" || *data == "" {
			log.Fatal("Topic or data missing")
		}
		log.Print("Publishing to topic: " + *topic + "  - with data: " + *data)
		nc.Publish(*topic, []byte(*data))
		log.Print("Publishing press enter to finish")
		reader := bufio.NewReader(os.Stdin)
		cont, _ := reader.ReadString('\n')
		fmt.Print(cont)
	} else if *runType == "subscribe" {
		if *topic == "" {
			log.Fatal("Topic must be supplied to subscribe")
		}
		if *queuGroup == "" {
			log.Print("Subscribing to topic: " + *topic)
			nc.Subscribe(*topic, func(m *nats.Msg) {
				log.Print("Topic::" + *topic + "   message::" + string(m.Data))
			})
		} else {
			log.Print("subscribing to topic: " + *topic + " under queue group: " + *queuGroup)
			nc.QueueSubscribe(*topic, *queuGroup, func(m *nats.Msg) {
				log.Print("Topic::" + *topic + "   message::" + string(m.Data))
			})
		}
		for {
		}
	}

}
