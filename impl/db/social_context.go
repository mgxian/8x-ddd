package db

import "ddd/models/user"

type Contact struct{}

func NewContact(user user.User) Contact {
	return Contact{}
}

func (c Contact) Make()                {}
func (c Contact) Break()               {}
func (c Contact) IsFriend(u user.User) {}

type SocialContextDB struct{}

func NewSocialContextDB() SocialContextDB {
	return SocialContextDB{}
}

func (s SocialContextDB) AsContact(u user.User) user.Contact {
	return NewContact(u)
}
