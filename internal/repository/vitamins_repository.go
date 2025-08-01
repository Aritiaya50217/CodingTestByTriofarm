package repository

import (
	"database/sql"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type VitaminRepository interface {
	CreateVitamin(vitamin *domain.Vitamins) error
	GetVitaminByName(name string) (*domain.Vitamins, error)
	GetMaxIndex() (int, error)
}

type vitaminRepository struct {
	db *gorm.DB
}

func NewVitaminRepository(db *gorm.DB) VitaminRepository {
	return &vitaminRepository{db: db}
}

func (r *vitaminRepository) CreateVitamin(vitamin *domain.Vitamins) error {
	return r.db.Create(vitamin).Error
}

func (r *vitaminRepository) GetVitaminByName(name string) (*domain.Vitamins, error) {
	var vitamin domain.Vitamins
	err := r.db.Where("name = ?", name).First(&vitamin).Error
	if err != nil {
		return nil, err
	}
	return &vitamin, nil
}

func (r *vitaminRepository) GetMaxIndex() (int, error) {
	var result sql.NullInt64

	err := r.db.Model(&domain.Vitamins{}).
		Select("MAX([index])").
		Scan(&result).Error

	if err != nil {
		return 0, err
	}

	if result.Valid {
		return int(result.Int64), nil
	}
	return 0, nil
}
