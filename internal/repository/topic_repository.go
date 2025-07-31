package repository

import (
	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/gorm"
)

type topicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) domain.TopicRepository {
	return &topicRepository{db: db}
}

func (r *topicRepository) GetAll() ([]domain.Topic, error) {
	var topics []domain.Topic
	err := r.db.Find(&topics).Error
	return topics, err
}
