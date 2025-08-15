package auth

import (
	"fmt"
	"log"
	"unicode"
)

type PasswordStrength struct {
	MinLength        int
	MaxLength        int
	RequireUppercase bool
	RequireLowercase bool
	RequireNumbers   bool
	RequireSpecial   bool
}

func DefaultPasswordStrength() PasswordStrength {
	return PasswordStrength{
		MinLength:        8,
		MaxLength:        72,
		RequireUppercase: true,
		RequireLowercase: true,
		RequireNumbers:   true,
		RequireSpecial:   true,
	}
}

func ValidatePasswordStrength(password string, strength PasswordStrength) []string {
	// NOSSA NOVA PISTA DE DEBUG:
	log.Printf("DEBUG: Validando a senha: '%s' (comprimento: %d)", password, len(password))
	var errors []string

	if len(password) < strength.MinLength {
		errors = append(errors, fmt.Sprintf("Password must be at least %d characters long", strength.MinLength))
	}
	if len(password) > strength.MaxLength {
		errors = append(errors, fmt.Sprintf("Password must be no more than %d characters long", strength.MaxLength))
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if strength.RequireUppercase && !hasUpper {
		errors = append(errors, "Password must contain at least one uppercase letter")
	}
	if strength.RequireLowercase && !hasLower {
		errors = append(errors, "Password must contain at least one lowercase letter")
	}
	if strength.RequireNumbers && !hasNumber {
		errors = append(errors, "Password must contain at least one number")
	}
	if strength.RequireSpecial && !hasSpecial {
		errors = append(errors, "Password must contain at least one special character")
	}

	return errors
}
