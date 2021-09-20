package db

import "ddd/models/user"

type UserRepositoryDB struct {
	db                  string
	subscriptionContext user.SubscriptionContext
	socialContext       user.SocialContext
	orderContext        user.OrderContext
}

func (r *UserRepositoryDB) GetUser(id string) user.User {
	u := user.User{}
	return u
}

func NewUserRepository(db string,
	subscriptionContext user.SubscriptionContext,
	socialContext user.SocialContext,
	orderContext user.OrderContext) *UserRepositoryDB {
	return &UserRepositoryDB{
		db:                  db,
		subscriptionContext: subscriptionContext,
		socialContext:       socialContext,
		orderContext:        orderContext,
	}
}

func (r *UserRepositoryDB) InSubscriptionContext() user.SubscriptionContext {
	return r.subscriptionContext
}

func (r *UserRepositoryDB) InSocialContext() user.SocialContext {
	return r.socialContext
}

func (r *UserRepositoryDB) InOrderContext() user.OrderContext {
	return r.orderContext
}
