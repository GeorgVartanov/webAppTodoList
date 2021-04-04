package adding

// RepositoryCreater ...
type Repository interface {
	AddUser(*User) error
}
// Service
type Service interface {
	AddUser(User) error
}

type service struct {
	repository Repository
}

// NewService adding
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser User Service
func (s *service) AddUser(user User) error {
	if err := s.repository.AddUser(&user); err != nil {
		return err
	}
	return nil
}
