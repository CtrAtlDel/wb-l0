package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"ivankvasov/publisher/utils"
	"log"
)

var natsUrl = "nats://localhost:4222"

func main() {
	model := utils.ReadConfig("/Users/ivankvasov/wb/l0/project/resourses/model.json")
	modelJson, _ := json.Marshal(model)
	fmt.Println("Publisher working...")
	sc, err := stan.Connect("test-cluster", "14", stan.NatsURL(natsUrl))
	if err != nil {
		log.Println(err)
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			log.Println(err)
		}
	}(sc)

	err = sc.Publish("my-channel", modelJson)

	if err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}
}
