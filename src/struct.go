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
	Name  string
	Types string
	Image string
}

func InitJeux() Jeux {
	return Jeux{}
}

type Check struct {
	Check_pseudo    bool
	Check_mdp       bool
	Check_connexion bool
}

func InitCheck() Check {
	return Check{}
}
