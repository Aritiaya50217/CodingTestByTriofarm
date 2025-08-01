package repository

import (
	"database/sql"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type MicroorganismRepository interface {
	CreateMicroorganism(microorganism *domain.Microorganisms) error
	GetMicroorganismByName(name string) (*domain.Microorganisms, error)
	GetMaxIndex() (int, error)
	GetAll() ([]domain.Microorganisms, error)
}

type microorganismsRepository struct {
	db *gorm.DB
}

func NewMicroorganismRepository(db *gorm.DB) MicroorganismRepository {
	return &microorganismsRepository{db: db}
}

func (r *microorganismsRepository) CreateMicroorganism(microorganism *domain.Microorganisms) error {
	return r.db.Create(microorganism).Error
}

func (r *microorganismsRepository) GetMicroorganismByName(name string) (*domain.Microorganisms, error) {
	var microorganism domain.Microorganisms

	err := r.db.Where("name = ?", name).First(&microorganism).Error
	if err != nil {
		return nil, err
	}
	return &microorganism, nil
}

func (r *microorganismsRepository) GetMaxIndex() (int, error) {
	var result sql.NullInt64
	err := r.db.Model(&domain.Microorganisms{}).
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

func (r *microorganismsRepository) GetAll() ([]domain.Microorganisms, error) {
	var microorganisms []domain.Microorganisms

	err := r.db.Model(&domain.Microorganisms{}).
		Select("id, name, topic_id,[index]").
		Order("[index] ASC").
		Find(&microorganisms).Error
	if err != nil {
		return nil, err
	}

	return microorganisms, nil
}
