package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	m "ivankvasov/project/internal/model"
	"ivankvasov/project/internal/service"
	"net/http"
)

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
