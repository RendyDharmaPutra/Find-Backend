package user

type User struct {
	ID string `gorm:"type:char(36);primaryKey"`
	Fullname string `gorm:"type:varchar(50);not null"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}