package person

type service struct {
	repo Repository
}

type Service interface {
	GetAllPersons() ([]Person, error)
	CreatePerson(person *Person) error
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (service *service) GetAllPersons() ([]Person, error) {
	persons, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (service *service) CreatePerson(person *Person) error {
	if err := service.repo.SavePerson(person); err != nil {
		return err
	}

	return nil
}