package service

import (
	"ivankvasov/project/internal/config"
	m "ivankvasov/project/internal/model"
	"ivankvasov/project/internal/repository"
)

func GetModelById(id string) *m.Model {
	model, ok := config.FoundModelInCacheById(id)
	if ok {
		return model
	}
	return repository.GetModelById(id)
}
