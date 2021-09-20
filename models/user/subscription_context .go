package user

import (
	"ddd/models/column"
)

type Reader interface {
	CanView()
	GetMySubscription() MySubscription
	Transfer(c column.Column, friend User)
}

type SubscriptionContext interface {
	AsReader(user User) Reader
}
