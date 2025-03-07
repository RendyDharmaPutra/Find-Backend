package features_init

import (
	"Find-Backend/features/person"
	"Find-Backend/features/user"

	"gorm.io/gorm"
)

type Repositories struct {
	UserRepo user.Repository
	PersonRepo person.Repository
}

func InitializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepo: user.NewRepository(db),
		PersonRepo: person.NewRepository(db),
	}
}