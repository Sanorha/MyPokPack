package pokemon

type User struct {
	Pseudo     string
	Motdepasse string
	Email      string
}
type Jeux struct {
	Name  string `json:"name"`
	Types string `json:"types"`
}

func InitUser() User {
	return User{}
}
func InitJeux() Jeux {
	return Jeux{}
}
