package usecase

import "github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"

type TopicUsecase struct {
	repo domain.TopicRepository
}

func NewTopicUsecase(repo domain.TopicRepository) *TopicUsecase {
	return &TopicUsecase{repo: repo}
}

func (uc *TopicUsecase) ListTopics() ([]domain.Topic, error) {
	return uc.repo.GetAll()
}
