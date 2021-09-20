package main

import (
	"ddd/impl/db"
	"ddd/models/column"
)

func main() {
	repo := db.NewUserRepository("mysql",
		db.NewSubscriptionContextDB(),
		db.NewSocialContextDB(),
		db.NewOrderContextDB())
	u := repo.GetUser("1")

	reader := repo.InSubscriptionContext().AsReader(u)
	reader.GetMySubscription().SubList()

	friend := repo.GetUser("2")
	c := column.NewColumn()
	reader.Transfer(c, friend)
}
