package models

type Category struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Priority  uint   `json:"priority"`
	Notify    bool   `json:"notify"`
	Recurrent bool   `json:"recurrent"`
	Name      string `gorm:"not null" json:"name"`
	UserID    uint   `json:"user_id"`
}
