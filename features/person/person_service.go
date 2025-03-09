package person

type service struct {
	repo Repository
}

type Service interface {
	GetAllPersons(userID interface{}) ([]Person, error)
	CreatePerson(person *Person, userID interface{}) error
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (service *service) GetAllPersons(userID interface{}) ([]Person, error) {
	persons, err := service.repo.FindAll(map[string]interface{}{"user_id": userID})
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (service *service) CreatePerson(person *Person, userID interface{}) error {
	person.UserID = userID.(string)
	
	if err := service.repo.SavePerson(person); err != nil {
		return err
	}

	return nil
}