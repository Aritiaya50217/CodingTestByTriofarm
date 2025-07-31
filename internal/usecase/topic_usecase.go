package usecase

import (
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
)

type TopicUsecase struct {
	repo repository.TopicRepository
}

func NewTopicUsecase(repo repository.TopicRepository) *TopicUsecase {
	return &TopicUsecase{repo: repo}
}

func (uc *TopicUsecase) ListTopics() ([]domain.Topic, error) {
	return uc.repo.GetAll()
}
