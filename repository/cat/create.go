package cat

import "github.com/viniosilva/go-api/model"

func (repo *catRepository) Create(cat *model.Cat) *model.Cat {
	repo.gormDB.Create(cat)

	return cat
}
