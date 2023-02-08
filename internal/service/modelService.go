package service

import (
	"ivankvasov/project/internal/config"
	m "ivankvasov/project/internal/model"
	"ivankvasov/project/internal/repository"
)

func InsertModel(m *m.Model) {
	config.InsertModelInCache(m)
	repository.InsertModel(m)
}

func GetModelById(id string) *m.Model {
	model, ok := config.FoundModelInCacheById(id)
	if ok {
		return model
	}
	modelFromDb := repository.GetModelById(id)
	if modelFromDb == nil {
		return nil
	}

	return modelFromDb
}
