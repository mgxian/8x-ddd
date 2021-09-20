package db

import (
	"ddd/models/column"
	"ddd/models/user"
)

type Reader struct {
	user           user.User
	mySubscription user.MySubscription
	socialContext  user.SocialContext
}

func NewReader(u user.User) Reader {
	r := Reader{user: u}
	return r
}

func (r Reader) GetMySubscription() user.MySubscription {
	return r.mySubscription
}

func (r Reader) SetMySubscription(mySubscription user.MySubscription) {
	r.mySubscription = mySubscription
}

func (r Reader) CanView() {}

func (r Reader) Transfer(c column.Column, u user.User) {
	contact := r.socialContext.AsContact(r.user)
	contact.IsFriend(u)
}

type SubscriptionContextDB struct{}

func NewSubscriptionContextDB() *SubscriptionContextDB {
	return &SubscriptionContextDB{}
}

func (s *SubscriptionContextDB) AsReader(u user.User) user.Reader {
	r := NewReader(u)
	r.SetMySubscription(NewMySubscriptionDB())
	r.socialContext = NewSocialContextDB()
	return r
}
