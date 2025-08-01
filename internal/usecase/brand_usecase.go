package usecase

import (
	"errors"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/utils"
)

type BrandUsecase interface {
	CreateBrand(brand *domain.Brands) error
	GetMaxIndex() (int, error)
	GetAll() ([]domain.Brands, error)
	UpdateBrand(brand *domain.Brands) error
	GetBrandByID(id int) (*domain.Brands, error)
	DeleteBrand(id int) error
	SwapBrand(brand []domain.Brands) error
}

type brandUsecase struct {
	brandRepo repository.BrandRepository
}

func NewBrandUsecase(repo repository.BrandRepository) BrandUsecase {
	return &brandUsecase{brandRepo: repo}
}

func (u brandUsecase) CreateBrand(brand *domain.Brands) error {
	// check name
	existingBrand, _ := u.brandRepo.GetBrandByName(brand.Name)
	if existingBrand != nil {
		return errors.New("name already exists")
	}

	return u.brandRepo.CreateBrand(brand)
}

func (u brandUsecase) GetMaxIndex() (int, error) {
	return u.brandRepo.GetMaxIndex()
}

func (u brandUsecase) GetAll() ([]domain.Brands, error) {
	return u.brandRepo.GetAll()
}

func (u brandUsecase) UpdateBrand(brand *domain.Brands) error {
	// check name
	existingBrand, err := u.brandRepo.GetBrandByName(brand.Name)
	if err == nil && existingBrand != nil {
		return errors.New("name already exists")
	}
	return u.brandRepo.UpdateBrand(brand)
}

func (u brandUsecase) GetBrandByID(id int) (*domain.Brands, error) {
	return u.brandRepo.GetBrandByID(id)
}

func (u brandUsecase) DeleteBrand(id int) error {
	return u.brandRepo.DeleteBrand(id)
}

func (u brandUsecase) SwapBrand(brands []domain.Brands) error {
	indexs := []int{}
	for _, brand := range brands {
		indexs = append(indexs, brand.Index)
	}

	if utils.IsDuplicateIndex(indexs) {
		return errors.New("duplicate index values found")
	}

	return u.brandRepo.SwapBrand(brands)
}
