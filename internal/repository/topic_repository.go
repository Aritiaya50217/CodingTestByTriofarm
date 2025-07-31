package repository

import (
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type TopicRepository interface {
	GetAll() ([]domain.Topic, error)
	// CreateTopic(name string) (domain.Topic, error)
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

// func (r *topicRepository) CreateTopic(name string) (*domain.Topic, error) {
// 	return r.db.Craete(topic).Error
// }
