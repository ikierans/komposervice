package types

type IUsers interface {
	GetByEmail(email string) (*Users, error)
	Create(_usr *Users) error
	Empty() bool
}
