package user

import (
	"errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() ([]User, error)
	FindOne(filters map[string]interface{}) (*User, error)
	SaveUser(user *User) error
	UpdateUser(id uint, value map[string]interface{}) error
}


func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) FindAll() ([]User, error) {
	var users []User
	
	if err := repo.db.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("pengguna tidak ditemukan")
		} else {
			log.Printf("Error tidak diketahui : %v", err.Error())

			err = errors.New("error tidak diketahui")
		}
		
		return nil, err
	}

	return users, nil
}

func (repo *repository) FindOne(filters map[string]interface{}) (*User, error) {	
	query := repo.db
	for key, value := range filters {
		query = query.Where(key +" = ?", value)
	}
	
	var user User
	if err := repo.db.First(&user, query).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("pengguna tidak ditemukan")
		} else {
			log.Printf("Error tidak diketahui : %v", err.Error())

			err = errors.New("error tidak diketahui")
		}
		
		return nil, err
	}

	return &user, nil
} 

func (repo *repository) SaveUser(user *User) error {
	if err := repo.db.Create(user).Error; err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = errors.New("username telah digunakan")
		} else {
			log.Printf("Error tidak diketahui : %v", err.Error())
			
			err =  errors.New("error tidak diketahui")
		}

		return err
	}

	return nil
}

func (repo *repository) UpdateUser(id uint, value map[string]interface{}) error {
	result :=  repo.db.Model(&User{}).Where("id = ?", id).Updates(value)
	if result.RowsAffected == 0 {
		return errors.New("pengguna tidak ditemukan")
	}

	if result.Error != nil {
		log.Printf("Error tidak diketahui : %v", result.Error.Error())

		return errors.New("error tidak diketahui")
	}

	return nil
}