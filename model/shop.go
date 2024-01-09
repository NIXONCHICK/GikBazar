package model

type Item struct {
	ID uint `json:"Item_id" gorm:"primaryKey"`

	Title       string `json:"Item_title" gorm:"not null;column:title;size:255"`
	Description string `json:"Description" gorm:"not null;column:description;size:255"`
	Coast       uint   `json:"Coast" gorm:"not null;column:coast"`
}
