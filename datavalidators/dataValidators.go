package datavalidators

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	// Simple regex for validating email format
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func IsValidUsername(password string) bool {
    // Check for potential SQL injection patterns
    regex := `^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?]*$`
    re := regexp.MustCompile(regex)
    return re.MatchString(password) && len(password) >= 8
}

func IsValidName(name string) bool {
    // Check for potential SQL injection patterns
    regex := `^[a-zA-Z]*$`
    re := regexp.MustCompile(regex)
    return re.MatchString(name)
}