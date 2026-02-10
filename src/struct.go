package pokemon

type User struct {
	Pseudo     string
	Motdepasse string
	Email      string
}

func InitUser() User {
	return User{}
}

type Jeux struct {
	Name  string `json:"name"`
	Types string `json:"types"`
}

func InitJeux() Jeux {
	return Jeux{}
}

type Check struct {
	Check_pseudo bool
	Check_mdp    bool
}

func InitCheck() Check {
	return Check{}
}
