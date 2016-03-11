package domain

const (
	SYSTEM_USER = 65535
)

type User struct {
	UserID uint
}

func CreatUser(id uint) *User {
	var user User
	user.UserID = id
	return &user
}
