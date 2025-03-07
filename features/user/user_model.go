package user

import "Find-Backend/features/person"

type User struct {
	ID string `gorm:"type:char(36);primaryKey" json:"id"`
	Fullname string `gorm:"type:varchar(50);not null" json:"fullname"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`

	Person []person.Person `gorm:"foreignKey:UserID" json:"persons"`
}