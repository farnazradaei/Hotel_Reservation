package user

type UserStorage interface {
	UserExsist(email string) (bool, error)
	GetUserByEmail(email string) (User, error)
	UserSave(user User) error
}