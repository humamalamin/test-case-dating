package conv

import (
	"crypto/sha1"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func StringToInt(s string) (int, error) {
	newData, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return newData, nil
}

func IntToString(i int) string {
	str := strconv.Itoa(i)
	return str
}

func HasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashShaPassword(password string, salt string) string {
	var saltedText = fmt.Sprintf("text: '%s', salt: '%s'", password, salt)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encripted = sha.Sum(nil)

	return fmt.Sprintf("%x", encripted)
}
