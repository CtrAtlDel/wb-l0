package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	c "ivankvasov/project/internal/cache"
	"ivankvasov/project/internal/controller"
	"ivankvasov/project/internal/model"
	m "ivankvasov/project/internal/model"
	s "ivankvasov/project/internal/service"
	"log"
	"net/http"
)

func init() {
	err := model.InitDb()
	if err == nil {
		c.InitCache()
	}
}

var natsUrl = "nats://localhost:4222"

func main() {
	defer m.Db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/models/{id}", controller.GetModelHandler).Methods("GET")
	router.HandleFunc("/model", controller.PostModelHandler).Methods("POST")
	http.Handle("/", router)

	sc, err := stan.Connect("test-cluster", "13", stan.NatsURL(natsUrl))
	if err != nil {
		log.Println(err)
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			log.Println(err)
		}
	}(sc)
	fmt.Println("Subscriber working...")
	if _, err := sc.Subscribe("my-channel", func(ms *stan.Msg) {
		var model m.Model

		err := json.Unmarshal(ms.Data, &model)

		if err != nil {
			fmt.Println("Incorrect data format...")
			return
		}
		s.InsertModel(&model)
		log.Printf("model: %s", string(ms.Data))
	}); err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}

	http.ListenAndServe(":8181", nil)
}
