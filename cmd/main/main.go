package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	c "ivankvasov/project/internal/config"
	"ivankvasov/project/internal/controller"
	"ivankvasov/project/internal/model"
	m "ivankvasov/project/internal/model"
	"log"
	"net/http"
)

func init() {
	err := model.InitDb()
	if err == nil {
		c.InitCache()
	}
}

var testId = "b563feb7b2b84b6test"

func main() {
	defer m.Db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/models/{id}", controller.GetModelHandler).Methods("GET")
	router.HandleFunc("/model", controller.PostModelHandler).Methods("POST")
	http.Handle("/", router)
	sc, err := stan.Connect("test-cluster", "13", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Println(err)
	}
	defer sc.Close()
	if _, err := sc.Subscribe("my-channel", func(m *stan.Msg) {
		log.Printf("Received message: %s", string(m.Data))
	}); err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}

	http.ListenAndServe(":8181", nil)
}
