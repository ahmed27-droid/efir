package models

type Level string

const (
	Normal   Level = "normal"
	Breaking Level = "breaking"
)

type Post struct {
	Base
	Name        string     `json:"name" gorm:"not null"`
	Broadcast   *Broadcast `json:"-" gorm:"foreignKey:BroadcastID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	BroadcastID uint       `json:"broadcast_id" gorm:"not null"`
	Level       Level      `json:"level" gorm:"type:varchar(20);check:level IN ('normal', 'breaking')"`
}
