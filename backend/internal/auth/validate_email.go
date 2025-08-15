package auth

import (
	"regexp"
	"strings"
)

func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}

	return matched
}

func ValidateEmailStrict(email string) []string {
	var errors []string

	if email == "" {
		errors = append(errors, "Email is required")
		return errors
	}

	if len(email) > 254 {
		errors = append(errors, "Email is too long (max 254 characters)")
	}

	if !strings.Contains(email, "@") {
		errors = append(errors, "Email must contain @ symbol")
		return errors
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		errors = append(errors, "Email must contain exactly one @ symbol")
		return errors
	}

	localPart := parts[0]
	domainPart := parts[1]

	if localPart == "" {
		errors = append(errors, "Email local part cannot be empty")
	} else {
		if len(localPart) > 64 {
			errors = append(errors, "Email local part is too long (max 64 characters)")
		}

		localRegex := `^[a-zA-Z0-9._%-]+$`
		if matched, _ := regexp.MatchString(localRegex, localPart); !matched {
			errors = append(errors, "Email local part contains invalid characters")
		}

		if strings.HasPrefix(localPart, ".") || strings.HasSuffix(localPart, ".") {
			errors = append(errors, "Email local part cannot start or end with a dot")
		}

		if strings.Contains(localPart, "..") {
			errors = append(errors, "Email local part cannot contain consecutive dots")
		}
	}

	if domainPart == "" {
		errors = append(errors, "Email domain cannot be empty")
	} else {
		if !strings.Contains(domainPart, ".") {
			errors = append(errors, "Email domain must contain at least one dot")
		}

		domainRegex := `^[a-zA-Z0-9.-]+$`
		if matched, _ := regexp.MatchString(domainRegex, domainPart); !matched {
			errors = append(errors, "Email domain contains invalid characters")
		}

		if strings.HasPrefix(domainPart, "-") || strings.HasSuffix(domainPart, "-") ||
			strings.HasPrefix(domainPart, ".") || strings.HasSuffix(domainPart, ".") {
			errors = append(errors, "Email domain cannot start or end with hyphen or dot")
		}

		domainParts := strings.Split(domainPart, ".")
		if len(domainParts) < 2 {
			errors = append(errors, "Email domain must have a valid TLD")
		} else {
			tld := domainParts[len(domainParts)-1]
			if len(tld) < 2 {
				errors = append(errors, "Email TLD must be at least 2 characters long")
			}

			tldRegex := `^[a-zA-Z]{2,}$`
			if matched, _ := regexp.MatchString(tldRegex, tld); !matched {
				errors = append(errors, "Email TLD must contain only letters")
			}
		}
	}

	if len(errors) == 0 {
		fullEmailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		if matched, _ := regexp.MatchString(fullEmailRegex, email); !matched {
			errors = append(errors, "Invalid email format")
		}
	}

	return errors
}

// IsValidEmailDomain verifica se o domínio do email é de uma lista permitida
func IsValidEmailDomain(email string, allowedDomains []string) bool {
	if !ValidateEmail(email) {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domain := strings.ToLower(parts[1])

	for _, allowedDomain := range allowedDomains {
		if domain == strings.ToLower(allowedDomain) {
			return true
		}
	}

	return false
}

// IsDisposableEmail verifica se é um email temporário/descartável
func IsDisposableEmail(email string) bool {
	disposableDomains := []string{
		"10minutemail.com",
		"guerrillamail.com",
		"mailinator.com",
		"tempmail.org",
		"yopmail.com",
		"temp-mail.org",
		"throwaway.email",
		"dispostable.com",
		"maildrop.cc",
		"getnada.com",
		"trashmail.com",
		"spambox.me",
		"fakeinbox.com",
		"mailnesia.com",
		"mytemp.email",
		"temp-mail.com",
		"moakt.com",
		"sharklasers.com",
		"mailcatch.com",
		"tempinbox.com",
		"spamex.com",
		"spambog.com",
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	domain := strings.ToLower(parts[1])

	for _, disposableDomain := range disposableDomains {
		if domain == disposableDomain {
			return true
		}
	}

	return false
}
