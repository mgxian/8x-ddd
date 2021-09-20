package user

type Repository interface {
	GetUser(id string) User
}
