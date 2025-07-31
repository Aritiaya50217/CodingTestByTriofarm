package domain

import (
	"time"
)

const (
	Medicine       = "ยา"
	Vitamins       = "วิตามิน"
	Microorganisms = "จุลินทรีย์"
	Brands         = "ยี่ห้อ"
)

type Topic struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
