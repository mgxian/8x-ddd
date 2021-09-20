package user

type Repository interface {
	GetUser(id string) User

	InSubscriptionContext() SubscriptionContext
	InSocialContext() SocialContext
	InOrderContext() OrderContext
}
