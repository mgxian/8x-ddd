package db

import "ddd/models/user"

type Buyer struct{}

func NewBuyer(u user.User) Buyer {
	return Buyer{}
}

func (b Buyer) PlaceOrder() {}

type OrderContextDB struct{}

func NewOrderContextDB() OrderContextDB {
	return OrderContextDB{}
}

func (o OrderContextDB) AsBuyer(u user.User) user.Buyer {
	return NewBuyer(u)
}
