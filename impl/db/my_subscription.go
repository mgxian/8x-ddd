package db

type MySubscriptionDB struct {
}

func (m MySubscriptionDB) SubList() {
}

func NewMySubscriptionDB() *MySubscriptionDB {
	return &MySubscriptionDB{}
}
