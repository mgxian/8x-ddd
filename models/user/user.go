package user

type User struct {
	name           string
	mySubscription MySubscription
}

func (u *User) GetMySubscription() MySubscription {
	return u.mySubscription
}

func (u *User) SetMySubscription(mySubscription MySubscription) {
	u.mySubscription = mySubscription
}
