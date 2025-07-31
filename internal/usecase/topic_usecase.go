package usecase

import (
	"errors"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/repository"
)

type TopicUsecase interface {
	ListTopics() ([]domain.Topic, error)
	CreateTopic(user *domain.Topic) error
	DeleteTopic(id uint) error
}

type topicUsecase struct {
	topicRepo repository.TopicRepository
}

func NewTopicUsecase(repo repository.TopicRepository) TopicUsecase {
	return &topicUsecase{topicRepo: repo}
}

func (uc *topicUsecase) ListTopics() ([]domain.Topic, error) {
	return uc.topicRepo.GetAll()
}

func (uc *topicUsecase) CreateTopic(topic *domain.Topic) error {
	// check name
	existingTopic, _ := uc.topicRepo.GetTopicName(topic.Name)
	if existingTopic != nil {
		return errors.New("name already exists")
	}
	return uc.topicRepo.CreateTopic(topic)
}

func (us *topicUsecase) DeleteTopic(id uint) error {
	return us.topicRepo.DeleteTopic(id)
}
