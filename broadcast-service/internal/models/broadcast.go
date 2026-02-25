package models

import "time"

type BroadcastStatus string

const (
	Draft   BroadcastStatus = "draft"
	Planned BroadcastStatus = "planned"
	Live    BroadcastStatus = "live"
	Ended   BroadcastStatus = "ended"
)

type Broadcast struct {
	Base
	Category   *Category       `json:"-" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CategoryID uint            `json:"category_id" gorm:"not null"`
	Name       string          `json:"name" gorm:"type:varchar(255);not null"`
	Status     BroadcastStatus `json:"status" gorm:"type:varchar(20);check:status IN ('draft', 'planned', 'live', 'ended')"`
	StartedAt  *time.Time      `json:"started_at"`
}
