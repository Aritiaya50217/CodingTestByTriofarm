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
	GetAll() ([]domain.Medicines, error)
	UpdateMedicine(medicine *domain.Medicines) error
	GetMedicineByName(name string) (*domain.Medicines, error)
	GetMedicineByID(id int) (*domain.Medicines, error)
	DeleteMedicine(id int) error
	SwapMedicines(medicines []domain.Medicines) error
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

func (r *medicineRepository) GetAll() ([]domain.Medicines, error) {
	var medicines []domain.Medicines
	err := r.db.Model(&domain.Medicines{}).
		Select("id, name, topic_id,[index]").
		Order("[index] ASC").
		Find(&medicines).Error
	if err != nil {
		return nil, err
	}
	return medicines, nil

}

func (r *medicineRepository) UpdateMedicine(medicine *domain.Medicines) error {
	err := r.db.Model(&domain.Medicines{}).Where("id = ? ", medicine.ID).Updates(domain.Medicines{Name: medicine.Name}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *medicineRepository) GetMedicineByName(name string) (*domain.Medicines, error) {
	var medicine domain.Medicines
	err := r.db.Where("name = ?", name).First(&medicine).Error
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *medicineRepository) GetMedicineByID(id int) (*domain.Medicines, error) {
	var medicine domain.Medicines
	err := r.db.Preload("Topic").Where("id = ?", id).First(&medicine).Error
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *medicineRepository) DeleteMedicine(id int) error {
	return r.db.Delete(&domain.Medicines{}, id).Error
}

func (r *medicineRepository) SwapMedicines(medicines []domain.Medicines) error {
	tx := r.db.Begin()

	for _, medicine := range medicines {
		if err := tx.Model(&domain.Medicines{}).
			Where("id = ?", medicine.ID).
			Updates(domain.Medicines{
				Index: medicine.Index,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}

	}
	return tx.Commit().Error
}
