package usecase

import (
	"errors"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/utils"
)

type MicroorganismUsecase interface {
	CreateMicroorganism(microorganism *domain.Microorganisms) error
	GetMaxIndex() (int, error)
	GetAllMicroorganism() ([]domain.Microorganisms, error)
	UpdateMicroorganism(microorganism *domain.Microorganisms) error
	GetMicroorganismByID(id int) (*domain.Microorganisms, error)
	DeleteMicroorganism(id int) error
	SwapMicroorganisms(microorganism []domain.Microorganisms) error
}

type microorganismUsecase struct {
	microorganismRepo repository.MicroorganismRepository
}

func NewMicroorganismUsecase(repo repository.MicroorganismRepository) MicroorganismUsecase {
	return &microorganismUsecase{microorganismRepo: repo}
}

func (u *microorganismUsecase) CreateMicroorganism(microorganism *domain.Microorganisms) error {
	// check name
	existingMicroorganism, _ := u.microorganismRepo.GetMicroorganismByName(microorganism.Name)
	if existingMicroorganism != nil {
		return errors.New("name already exists")
	}

	return u.microorganismRepo.CreateMicroorganism(microorganism)
}

func (u *microorganismUsecase) GetMaxIndex() (int, error) {
	return u.microorganismRepo.GetMaxIndex()
}

func (u *microorganismUsecase) GetAllMicroorganism() ([]domain.Microorganisms, error) {
	return u.microorganismRepo.GetAll()
}

func (u *microorganismUsecase) UpdateMicroorganism(microorganism *domain.Microorganisms) error {
	// check name
	existingMicroorganisms, err := u.microorganismRepo.GetMicroorganismByName(microorganism.Name)
	if err == nil && existingMicroorganisms != nil {
		return errors.New("name already exists")
	}
	return u.microorganismRepo.UpdateMicroorganism(microorganism)
}

func (u *microorganismUsecase) GetMicroorganismByID(id int) (*domain.Microorganisms, error) {
	return u.microorganismRepo.GetMicroorganismByID(id)
}

func (u *microorganismUsecase) DeleteMicroorganism(id int) error {
	return u.microorganismRepo.DeleteMicroorganism(id)
}

func (u *microorganismUsecase) SwapMicroorganisms(microorganisms []domain.Microorganisms) error {
	indexs := []int{}
	for _, microorganism := range microorganisms {
		indexs = append(indexs, microorganism.Index)
	}

	if utils.IsDuplicateIndex(indexs) {
		return errors.New("duplicate index values found")
	}

	return u.microorganismRepo.SwapMicroorganisms(microorganisms)
}
