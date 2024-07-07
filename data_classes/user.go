package data_classes

type User struct {
	FirstName string
	LastName string
	Email string
}

func (user User) GetFullName() string {
	return user.FirstName + " " + user.LastName
}