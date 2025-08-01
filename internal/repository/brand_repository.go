package repository

import (
	"database/sql"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type BrandRepository interface {
	CreateBrand(brand *domain.Brands) error
	GetBrandByName(name string) (*domain.Brands, error)
	GetMaxIndex() (int, error)
	GetAll() ([]domain.Brands, error)
	UpdateBrand(brand *domain.Brands) error
	GetBrandByID(id int) (*domain.Brands, error)
	DeleteBrand(id int) error
	SwapBrand(brand []domain.Brands) error
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}

func (r *brandRepository) CreateBrand(brand *domain.Brands) error {
	return r.db.Create(brand).Error
}

func (r *brandRepository) GetBrandByName(name string) (*domain.Brands, error) {
	var brand domain.Brands

	err := r.db.Where("name = ?", name).First(&brand).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (r *brandRepository) GetMaxIndex() (int, error) {
	var result sql.NullInt64
	err := r.db.Model(&domain.Brands{}).
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

func (r *brandRepository) GetAll() ([]domain.Brands, error) {
	var brands []domain.Brands

	err := r.db.Model(&domain.Brands{}).
		Select("id, name, topic_id,[index]").
		Order("[index] ASC").
		Find(&brands).Error
	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (r *brandRepository) UpdateBrand(brand *domain.Brands) error {
	err := r.db.Model(&domain.Brands{}).
		Where("id = ? ", brand.ID).
		Updates(domain.Brands{Name: brand.Name}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *brandRepository) GetBrandByID(id int) (*domain.Brands, error) {
	var brand domain.Brands
	err := r.db.Preload("Topic").Where("id = ?", id).First(&brand).Error
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (r *brandRepository) DeleteBrand(id int) error {
	return r.db.Delete(&domain.Brands{}, id).Error
}

func (r *brandRepository) SwapBrand(brands []domain.Brands) error {
	tx := r.db.Begin()

	for _, brand := range brands {
		if err := tx.Model(&domain.Brands{}).
			Where("id = ?", brand.ID).
			Updates(domain.Brands{
				Index: brand.Index,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}

	}
	return tx.Commit().Error
}
