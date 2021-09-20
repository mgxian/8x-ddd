package user

type Contact interface {
	Make()
	Break()
	IsFriend(u User)
}

type SocialContext interface {
	AsContact(user User) Contact
}
