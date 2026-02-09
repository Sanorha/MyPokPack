package pokemon

import (
	"golang.org/x/crypto/bcrypt"
)

const pepper = "C:h6_yzI-2/$zjé)dW,;!zk84gPù%"

func CreateHashMDP(password string) (string, error) {
	password_pepper := password + pepper

	bytes, err := bcrypt.GenerateFromPassword([]byte(password_pepper), 14)

	return string(bytes), err
}

func CompareMDP(password string) {
	mdp_baseSQL := []byte("")
	///////////////////////////////////////////////////////
	// REQUETE SQL POUR recup MDP sur base donnée		 //
	///////////////////////////////////////////////////////

	password_bytes := []byte(password)
	hashedPassword := []byte(mdp_baseSQL)

	bcrypt.CompareHashAndPassword(hashedPassword, password_bytes)

}
