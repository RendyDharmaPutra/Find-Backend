package person

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() ([]Person, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) FindAll() ([]Person, error) {
	var persons []Person

	if err := repo.db.Find(&persons).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("pengguna tidak ditemukan")
		} else {
			log.Printf("Error tidak diketahui : %v", err.Error())

			err = errors.New("error tidak diketahui")
		}
		
		return nil, err
	}

	return persons, nil
}