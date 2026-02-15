package pokemon

type User struct {
	Pseudo     string
	Motdepasse string
	Email      string
}

func InitUser() User {
	return User{}
}

type Jeux_slice struct {
	Jeux_slice []Jeux
}

type Jeux struct {
	Name  string
	Types string
	Image string
}

func InitJeux() Jeux_slice {
	return Jeux_slice{}
}

type Check struct {
	Check_pseudo    bool
	Check_mdp       bool
	Check_connexion bool
}

func InitCheck() Check {
	return Check{}
}

type Booster struct {
	Booster []Pokemon
}

type Pokemon struct {
	Name  string
	Types string
	Image string
}

func InitBooster() Booster {
	return Booster{}
}
