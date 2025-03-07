package person

type service struct {
	repo Repository
}

type Service interface {
	GetAllPersons() ([]Person, error)
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