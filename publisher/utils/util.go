package utils

import (
	"encoding/json"
	"fmt"
	m "ivankvasov/publisher/model"
	"os"
)

func ReadConfig(s string) *m.Model {
	var model m.Model
	jsn, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("Error from open file:", err)
		return nil
	}

	err = json.Unmarshal(jsn, &model)
	if err != nil {
		fmt.Println("Error from open file:", err)
		return nil
	}

	return &model
}
