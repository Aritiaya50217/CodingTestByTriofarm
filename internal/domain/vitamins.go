package domain

import "time"

type Vitamins struct {
	ID        int       `gorm:"primaryKey" json:"id" `
	Name      string    `json:"name"`
	TopicID   uint      `json:"topic_id"` // Foreign Key
	Topic     Topic     `gorm:"foreignKey:TopicID"`
	Index     int       `gorm:"default:1" json:"index"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
