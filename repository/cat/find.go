package cat

import "github.com/viniosilva/go-api/model"

func (repo *catRepository) Find() []model.Cat {
	var cats []model.Cat
	repo.gormDB.Find(&cats)

	return cats
}
