package service

import (
	"ivankvasov/project/internal/cache"
	m "ivankvasov/project/internal/model"
	"ivankvasov/project/internal/repository"
)

func InsertModel(m *m.Model) {
	cache.InsertModelInCache(m)
	repository.InsertModel(m)
}

func GetModelById(id string) *m.Model {
	model, ok := cache.FoundModelInCacheById(id)
	if ok {
		return model
	}
	modelFromDb := repository.GetModelById(id)
	if modelFromDb == nil {
		return nil
	}

	return modelFromDb
}
