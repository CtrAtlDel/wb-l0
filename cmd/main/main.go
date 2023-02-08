package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	c "ivankvasov/project/internal/config"
	"ivankvasov/project/internal/controller"
	"ivankvasov/project/internal/model"
	m "ivankvasov/project/internal/model"
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

	http.ListenAndServe(":8181", nil)
}
