package config

import (
	m "ivankvasov/project/internal/model"
	"log"
)

type Cache struct {
	items map[string]m.Model
}

var cache Cache

func InitCache() {
	cache.items = make(map[string]m.Model)
	rows, err := m.Db.Query("SELECT o.order FROM orders as o")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var model m.Model
		err := rows.Scan(&model)
		if err != nil {
			log.Println(err)
		}
		cache.items[model.OrderUId] = model
	}
}

func InsertModelInCache(m *m.Model) {
	cache.items[m.OrderUId] = *m
}

func FoundModelInCacheById(id string) (*m.Model, bool) {
	model, ok := cache.items[id]

	return &model, ok
}
