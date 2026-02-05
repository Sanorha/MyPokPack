package pokemon

type User struct {
	Pseudo     string
	Motdepasse string
	Email      string
}

func InitUser() User {
	return User{}
}
