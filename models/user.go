package models

type User struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Name     string `gorm:"not null" json:"name"`
    Email    string `gorm:"not null" json:"email"`
    Username string `json:"username"`
    Password string `gorm:"not null" json:"password,omitempty"`
}
