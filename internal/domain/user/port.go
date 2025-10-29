package user

type UserRepository interface {
	Save(*User) (*User, error)
	FindByID(id int) (*User, error)
}