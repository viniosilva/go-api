package repository

import (
	"github.com/viniosilva/go-api/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=cat_repository_mock.go -package=repository . CatRepository
type CatRepository interface {
	Create(cat *model.Cat) *model.Cat
	Find() []model.Cat
}

type catRepositoryImpl struct {
	gormDB *gorm.DB
}

func NewCatRepository(gormDB *gorm.DB) CatRepository {
	return &catRepositoryImpl{gormDB: gormDB}
}

func (impl catRepositoryImpl) Create(cat *model.Cat) *model.Cat {
	impl.gormDB.Create(cat)

	return cat
}

func (impl catRepositoryImpl) Find() []model.Cat {
	var cats []model.Cat
	impl.gormDB.Find(&cats)

	return cats
}
