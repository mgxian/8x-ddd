package db

import "ddd/models/user"

type UserRepositoryDB struct {
	db string
}

func (r *UserRepositoryDB) GetUser(id string) user.User {
	u := user.User{}
	u.SetMySubscription(NewMySubscriptionDB())
	return u
}

func NewUserRepository(db string) *UserRepositoryDB {
	return &UserRepositoryDB{
		db: db,
	}
}
