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
	Name string `json:"name"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`

	Sprites struct {
		Other struct {
			Dream_world struct {
				Image string `json:"front_default"`
			} `json:"dream_world"`
		} `json:"other"`
	} `json:"sprites"`
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
