package domain

import "github.com/jmoiron/sqlx"

type UserRepositoryDb struct {
	client *sqlx.DB
}

func (u UserRepositoryDb) FindById(id string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryDb) FindAll() ([]User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepositoryDb() UserRepositoryDb {

	return UserRepositoryDb{client: nil}

}
