package repository

import (
	"encoding/json"
	m "ivankvasov/project/internal/model"
	"log"
)

func InsertModel(model *m.Model) {
	modelJson, _ := json.Marshal(model)
	_, err := m.Db.Exec(`INSERT INTO public.orders ("order") VALUES ($1)`, modelJson)
	if err != nil {
		log.Println(err)
	}
}

func GetModelById(id string) *m.Model {
	var model m.Model
	err := m.Db.QueryRow("SELECT o.order FROM orders as o WHERE o.order ->> 'order_uid'= $1", &id).Scan(&model)
	if err != nil {
		log.Println(err)
	}
	return &model
}
