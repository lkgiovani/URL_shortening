package postgres

import (
	"url_shortening/environment"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(config *environment.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DB.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
