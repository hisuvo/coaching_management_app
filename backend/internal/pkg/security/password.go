package security

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plain-text password.
// This function can be reused anywhere in the application.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

// CheckPassword verifies a plain-text password against a stored hash.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}