package storage

import (
	"gorm.io/gorm"

	"github.com/kavladst/ads_api/internal/app/configuration"
)

type Storage struct {
	DB *gorm.DB
}

func New(config *configuration.Configuration) (*Storage, error) {
	db, err := InitDB(config.DBHost, config.DBPort, config.DBName, config.DBUser, config.DBPassword)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}
