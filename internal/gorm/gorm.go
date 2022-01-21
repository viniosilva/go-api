package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gorm interface {
	GetDB() *gorm.DB
	Migrate(models ...interface{})
}

type gormImpl struct {
	DB *gorm.DB
}

type GormParams struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func NewGorm(p GormParams) Gorm {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", p.Username, p.Password, p.Host, p.Port, p.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("DB -> %s", dsn)
		panic("failed to connect database")
	}

	return &gormImpl{DB: db}
}

func (impl *gormImpl) GetDB() *gorm.DB {
	return impl.DB
}

func (impl *gormImpl) Migrate(models ...interface{}) {
	impl.DB.AutoMigrate(models...)
}
