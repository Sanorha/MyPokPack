package pokemon

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const pepper = "C:h6_yzI-2/$zjé)dW,;!zk84gPù%"

func CreateHashMDP(password string) (string, error) {
	password_pepper := password + pepper

	bytes, err := bcrypt.GenerateFromPassword([]byte(password_pepper), 14)

	return string(bytes), err
}

func CompareMDP(motdepasse string, motdepasse_sql string) bool {
	motdepasse += pepper

	err := bcrypt.CompareHashAndPassword([]byte(motdepasse_sql), []byte(motdepasse))

	fmt.Println("mdp compare : ", err, motdepasse, motdepasse_sql)

	if err == nil {
		return true
	} else {
		return false
	}
}
