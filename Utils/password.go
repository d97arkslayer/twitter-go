package Utils

import "golang.org/x/crypto/bcrypt"

/**
 * EncryptPassword
 * This function encrypt the plain text password
 */
func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass),cost)
	return string(bytes), err
}

/**
 * ComparePassword
 * Use to compare provided password with database password
 */
func ComparePassword(password string, encryptedPassword string) bool{
	passwordBytes := []byte(password)
	databasePassword := []byte(encryptedPassword)
	err := bcrypt.CompareHashAndPassword(databasePassword, passwordBytes)
	if err != nil {
		return false
	}
	return true
}
