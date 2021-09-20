package main

import (
	"ddd/impl/db"
)

func main() {
	repo := db.NewUserRepository("mysql")
	u := repo.GetUser("1")
	u.GetMySubscription().SubList()
}
