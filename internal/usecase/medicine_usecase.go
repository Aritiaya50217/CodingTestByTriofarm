package usecase

import (
	"errors"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
)

type MedicineUsecase interface {
	CreateMedicine(medicines *domain.Medicines) error
	PreloadTopic(medicine *domain.Medicines) error
	GetMaxIndex() (int, error)
	GetAllMedicines() ([]domain.Medicines, error)
	UpdateMedicine(medicine *domain.Medicines) error
	GetMedicineByID(id int) (*domain.Medicines, error)
	DeleteMedicine(id int) error
}

type medicineUsecase struct {
	medicineRepo repository.MedicineRepository
}

func NewMedicineUsecase(repo repository.MedicineRepository) MedicineUsecase {
	return &medicineUsecase{medicineRepo: repo}
}

func (u *medicineUsecase) CreateMedicine(medicine *domain.Medicines) error {
	// check name
	existingMedicine, _ := u.medicineRepo.GetMedicineName(medicine.Name)
	if existingMedicine != nil {
		return errors.New("name already exists")
	}

	return u.medicineRepo.CreateMedicine(medicine)
}

func (u *medicineUsecase) PreloadTopic(medicine *domain.Medicines) error {
	return u.medicineRepo.PreloadTopic(medicine)
}

func (u *medicineUsecase) GetMaxIndex() (int, error) {
	return u.medicineRepo.GetMaxIndex()
}

func (u *medicineUsecase) GetAllMedicines() ([]domain.Medicines, error) {
	return u.medicineRepo.GetAll()
}

func (u *medicineUsecase) UpdateMedicine(medicine *domain.Medicines) error {
	// check name
	existingMedicine, err := u.medicineRepo.GetMedicineByName(medicine.Name)
	if err == nil && existingMedicine != nil {
		return errors.New("name already exists")
	}
	return u.medicineRepo.UpdateMedicine(medicine)
}

func (u *medicineUsecase) GetMedicineByID(id int) (*domain.Medicines, error) {
	return u.medicineRepo.GetMedicineByID(id)
}

func (u *medicineUsecase) DeleteMedicine(id int) error {
	return u.medicineRepo.DeleteMedicine(id)
}
