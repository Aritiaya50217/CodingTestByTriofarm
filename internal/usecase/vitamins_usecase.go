package usecase

import (
	"errors"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
)

type VitaminUsecase interface {
	CreateVitamin(vitamin *domain.Vitamins) error
	GetMaxIndex() (int, error)
}

type vitaminUsecase struct {
	vitaminRepo repository.VitaminRepository
}

func NewVitaminUsecase(repo repository.VitaminRepository) VitaminUsecase {
	return &vitaminUsecase{vitaminRepo: repo}
}

func (u *vitaminUsecase) CreateVitamin(vitamin *domain.Vitamins) error {
	// check name
	existingVitamin, _ := u.vitaminRepo.GetVitaminByName(vitamin.Name)
	if existingVitamin != nil {
		return errors.New("name already exists")
	}

	return u.vitaminRepo.CreateVitamin(vitamin)
}

func (u *vitaminUsecase) GetMaxIndex() (int, error) {
	return u.vitaminRepo.GetMaxIndex()
}
