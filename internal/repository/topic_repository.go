package repository

import (
	"fmt"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type TopicRepository interface {
	GetAll() ([]domain.Topic, error)
	CreateTopic(topic *domain.Topic) error
	GetTopicName(name string) (*domain.Topic, error)
	DeleteTopic(id uint) error
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

func (r *topicRepository) DeleteTopic(id uint) error {
	result := r.db.Delete(&domain.Topic{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("topic with ID %d not found", id)
	}
	return nil
}
