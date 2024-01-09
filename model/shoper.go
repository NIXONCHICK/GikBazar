package model

type Shoper struct {
	ID uint `json:"Shoper_id" gorm:"primaryKey"`

	UserId  uint `json:"UserId" gorm:"not null;column:user_id"`
	ItemsID uint `json:"ItemsId" gorm:"not null;column:items_id"`
	Amount  uint `json:"Amount" gorm:"not null;column:amount"`
}
