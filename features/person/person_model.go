package person

type Person struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Fullname string `gorm:"size:50;not null" json:"fullname"`
	Address string `gorm:"size:100" json:"address"`
	Job string `gorm:"size:50;not null" json:"job"`

	UserID string `gorm:"constraint:OnDelete:CASCADE" json:"user_id"`
}