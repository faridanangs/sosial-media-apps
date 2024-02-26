package util

import (
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pass string) string {
	pasGnrate, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		log.Fatal("Error generating password")
	}
	return string(pasGnrate)
}
