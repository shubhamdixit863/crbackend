package domain

type User struct {
	Name string
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindById(id string) (*User, error)
}
