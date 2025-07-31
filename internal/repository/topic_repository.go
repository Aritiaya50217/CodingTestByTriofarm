package repository

import (
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type TopicRepository interface {
	GetAll() ([]domain.Topic, error)
	CreateTopic(topic *domain.Topic) error
	GetTopicName(name string) (*domain.Topic, error)
}

type topicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	return &topicRepository{db: db}
}

func (r *topicRepository) GetAll() ([]domain.Topic, error) {
	var topics []domain.Topic
	err := r.db.Find(&topics).Error
	return topics, err
}

func (r *topicRepository) CreateTopic(topic *domain.Topic) error {
	return r.db.Create(&topic).Error
}

func (r *topicRepository) GetTopicName(name string) (*domain.Topic, error) {
	var topic domain.Topic
	err := r.db.Where("name = ?", name).First(&topic).Error
	if err != nil {
		return nil, err
	}

	return &topic, nil
}
