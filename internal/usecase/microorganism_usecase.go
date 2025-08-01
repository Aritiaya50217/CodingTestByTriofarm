package usecase

import (
	"errors"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
)

type MicroorganismUsecase interface {
	CreateMicroorganism(microorganism *domain.Microorganisms) error
	GetMaxIndex() (int, error)
	GetAllMicroorganism() ([]domain.Microorganisms, error)
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
