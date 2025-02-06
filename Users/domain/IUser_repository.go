package domain

type IUser interface {
	SaveUserFunction(user User) error
	GetUserFunction(id int32)([]User , error)
	EditUserFunction(user User)error
	DeleteUserFunction(id int32)error
	ComprobateQuantityOfUsers()(int, error)
}