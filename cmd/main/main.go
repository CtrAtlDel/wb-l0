package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io"
	c "ivankvasov/project/internal/config"
	"ivankvasov/project/internal/model"
	m "ivankvasov/project/internal/model"
	"ivankvasov/project/internal/service"
	"net/http"
)

func init() {
	err := model.InitDb()
	if err == nil {
		c.InitCache()
	}
}

var testId = "b563feb7b2b84b6test"

func GetModelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	model := service.GetModelById(id)

	modelJson, _ := json.Marshal(model)
	response := fmt.Sprintf("Product: %s", string(modelJson))
	fmt.Fprint(w, response)
}

func PostModelHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var model m.Model
	json.Unmarshal(reqBody, &model)
	service.InsertModel(&model)
}

func main() {
	defer m.Db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/models/{id}", GetModelHandler).Methods("GET")
	router.HandleFunc("/model", PostModelHandler).Methods("POST")
	http.Handle("/", router)

	http.ListenAndServe(":8181", nil)
}
