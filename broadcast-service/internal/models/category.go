package models

type Category struct {
	Base

	Name     string     `json:"name" gorm:"type:varchar(255);not null"`
	ParentID *uint      `json:"parent_id"`
	Parent   *Category  `json:"-" gorm:"foreignKey:ParentID"`
	Children []Category `json:"children" gorm:"foreignKey:ParentID"`
}
