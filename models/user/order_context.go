package user

type Buyer interface {
	PlaceOrder()
}

type OrderContext interface {
	AsBuyer(u User) Buyer
}
