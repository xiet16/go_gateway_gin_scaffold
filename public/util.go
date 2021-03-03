package public

import (
	"crypto/sha256"
	"fmt"
)

func GenSecretPassword(secret, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + secret))
	return fmt.Sprintf("%x", s2.Sum(nil))
}