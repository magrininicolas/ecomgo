package auth

import "golang.org/x/crypto/bcrypt"

func HashPasswd(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(dbPwd, rPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(rPwd))
	return err == nil
}
