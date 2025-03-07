package features_init

import (
	"Find-Backend/features/auth"
	"Find-Backend/features/person"
	"Find-Backend/features/user"

	"gorm.io/gorm"
)

type Module struct {
	UserService user.Service
	AuthService auth.Service
	PersonService person.Service
}

func InitializeModules(db *gorm.DB) *Module {
	repositories := InitializeRepositories(db)

	return &Module{
		UserService: user.NewService(repositories.UserRepo),
		AuthService: auth.NewService(repositories.UserRepo),
		PersonService: person.NewService(repositories.PersonRepo),
	}
}