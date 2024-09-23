package models

type Payment struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	CategoryID  uint    `json:"category_id"`
	NetAmount   float64 `gorm:"not null" json:"net_amount"`
	GrossAmount float64 `gorm:"not null" json:"gross_amount"`
	Deductible  float64 `json:"deductible"`
	Name        string  `gorm:"not null" json:"name"`
	ChargeDate  string  `gorm:"not null" json:"author"`
	Recurrent   bool    `json:"recurrent"`
	PaymentType string  `gorm:"not null" json:"payment_type"`
	Paid        bool    `gorm:"not null" json:"paid"`
}
