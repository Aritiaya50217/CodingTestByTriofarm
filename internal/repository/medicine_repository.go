package repository

import (
	"database/sql"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	CreateMedicine(medicine *domain.Medicines) error
	GetMedicineName(name string) (*domain.Medicines, error)
	PreloadTopic(medicine *domain.Medicines) error
	GetMaxIndex() (int, error)
}

type medicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) MedicineRepository {
	return &medicineRepository{db: db}
}

func (r *medicineRepository) CreateMedicine(medicine *domain.Medicines) error {
	return r.db.Create(medicine).Error
}

func (r *medicineRepository) GetMedicineName(name string) (*domain.Medicines, error) {
	var medicine domain.Medicines
	err := r.db.Where("name = ?", name).First(&medicine).Error
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *medicineRepository) PreloadTopic(medicine *domain.Medicines) error {
	return r.db.Preload("Topic").First(medicine, medicine.ID).Error
}

func (r *medicineRepository) GetMaxIndex() (int, error) {
	var result sql.NullInt64
	err := r.db.Model(&domain.Medicines{}).
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
