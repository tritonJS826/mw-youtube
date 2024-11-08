package service

type UserService struct {
	repository UserRepository
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"Name"`
}

func (s *UserService) GetUser(id string) (*User, error) {
	return s.repository.FindByID(id)
}
