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

func (uc *medicineUsecase) PreloadTopic(medicine *domain.Medicines) error {
	return uc.medicineRepo.PreloadTopic(medicine)
}

func (uc *medicineUsecase) GetMaxIndex() (int, error) {
	return uc.medicineRepo.GetMaxIndex()
}
