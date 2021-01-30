package storage

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ad struct {
	ID          uuid.UUID      `gorm:"type:primaryKey;uuid;default:uuid_generate_v4()"`
	Name        string         `gorm:"not null"`
	Description string         `gorm:"not null"`
	Price       int            `gorm:"not null"`
	PhotosUrls  pq.StringArray `gorm:"type:text[];not null"`
	CreatedAt   time.Time
}

func (s *Storage) GetAd(id uuid.UUID) (Ad, error) {
	var ad Ad
	err := s.DB.First(&ad, id).Error
	return ad, err
}

func (s *Storage) GetAds(page int, perPage int, sortedBy string, isReversed bool) ([]Ad, error) {
	if sortedBy != "price" && sortedBy != "created_at" {
		return nil, errors.New("sort by must be \"price\" or \"created_at\"")
	}
	if page <= 0 {
		return nil, errors.New("page must be positive int")
	}
	if perPage <= 0 {
		return nil, errors.New("per page must be positive int")
	}
	var orderString string
	if isReversed {
		orderString = fmt.Sprintf("%s desc", sortedBy)
	} else {
		orderString = sortedBy
	}
	ads := []Ad{}
	err := s.DB.Limit(perPage).Offset((page - 1) * perPage).Order(orderString).Model(&Ad{}).Find(&ads).Error
	if err == nil || errors.Is(err, gorm.ErrRecordNotFound) {
		return ads, nil
	}
	return nil, err
}

func (s *Storage) CreateAd(name string, description string, photoUrls []string, price int) (uuid.UUID, error) {
	ad := Ad{Name: name, Description: description, Price: price, PhotosUrls: photoUrls}
	err := s.DB.Create(&ad).Error
	return ad.ID, err
}
