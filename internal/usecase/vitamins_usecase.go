package usecase

import (
	"errors"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/utils"
)

type VitaminUsecase interface {
	CreateVitamin(vitamin *domain.Vitamins) error
	GetMaxIndex() (int, error)
	GetAllVitamin() ([]domain.Vitamins, error)
	UpdateVitamin(vitamin *domain.Vitamins) error
	GetVitaminByID(id int) (*domain.Vitamins, error)
	DeleteVitamin(id int) error
	SwapVitamins(vitamin []domain.Vitamins) error
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

func (u *vitaminUsecase) GetAllVitamin() ([]domain.Vitamins, error) {
	return u.vitaminRepo.GetAll()
}

func (u *vitaminUsecase) UpdateVitamin(vitamin *domain.Vitamins) error {
	// check name
	existingVitamin, err := u.vitaminRepo.GetVitaminByName(vitamin.Name)
	if err == nil && existingVitamin != nil {
		return errors.New("name already exists")
	}
	return u.vitaminRepo.UpdateVitamin(vitamin)
}

func (u *vitaminUsecase) GetVitaminByID(id int) (*domain.Vitamins, error) {
	return u.vitaminRepo.GetVitaminByID(id)
}

func (u *vitaminUsecase) DeleteVitamin(id int) error {
	return u.vitaminRepo.DeleteVitamin(id)
}

func (u *vitaminUsecase) SwapVitamins(vitamins []domain.Vitamins) error {
	indexs := []int{}
	for _, vitamin := range vitamins {
		indexs = append(indexs, vitamin.Index)
	}

	if utils.IsDuplicateIndex(indexs) {
		return errors.New("duplicate index values found")
	}
	return u.vitaminRepo.SwapVitamins(vitamins)
}
