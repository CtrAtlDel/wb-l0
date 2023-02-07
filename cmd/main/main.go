package main

import (
	_ "github.com/lib/pq"
	c "ivankvasov/project/internal/config"
	"ivankvasov/project/internal/model"
	m "ivankvasov/project/internal/model"
)

func init() {
	err := model.InitDb()
	if err == nil {
		c.InitCache()
	}
}

var testId = "b563feb7b2b84b6test"

func main() {
	defer m.Db.Close() // why this need ?
}
