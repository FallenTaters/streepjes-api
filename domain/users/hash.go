package users

import "golang.org/x/crypto/bcrypt"

func hash(password []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return hash
}
