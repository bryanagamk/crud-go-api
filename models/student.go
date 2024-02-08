package models

import "time"

type Student struct {
	IdStudent int64      `gorm:"primaryKey" json:"id_student"`
	Name      string     `gorm:"type:varchar(255)" json:"name"`
	Course    string     `gorm:"type:varchar(255)" json:"course"`
	Email     string     `gorm:"type:varchar(255)" json:"email"`
	Phone     string     `gorm:"type:varchar(255)" json:"phone"`
	CreatedAt time.Time  `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:datetime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:datetime" json:"deleted_at"`
}
