package validations

import (
	"regexp"

	"github.com/go-playground/validator"
)

var validate = validator.New()

// init registers the validation functions for the domain structs
func init() {
	validate.RegisterValidation("username", validateUsername)
	validate.RegisterValidation("password", validatePassword)
	validate.RegisterValidation("email", validateEmail)
}

func GetValidator() *validator.Validate {
	return validate
}

// validateUsername checks if the username is valid
// ruleset:
// - min 3 characters
// - max 20 characters
// - only letters, numbers, and underscores
// - must start with a letter
// - must be unique
// - cannot consecutively have the underscore
// - must be end with a letter or number
var (
	usernameRegex              = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	usernameStart              = regexp.MustCompile(`^[a-zA-Z]`)
	usernameEnd                = regexp.MustCompile(`[a-zA-Z0-9]$`)
	usernameNoDoubleUnderscore = regexp.MustCompile(`(_{2,})`)
	usernameMinLength          = 3
	usernameMaxLength          = 20
)

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	if len(username) < usernameMinLength || len(username) > usernameMaxLength {
		return false
	}

	if !usernameRegex.MatchString(username) {
		if !usernameStart.MatchString(username) {
			return false
		}

		if !usernameEnd.MatchString(username) {
			return false
		}
	}

	if usernameNoDoubleUnderscore.MatchString(username) {
		return false
	}

	return true
}

// validatePassword checks if the password is valid
// ruleset:
// - min 8 characters
// - max 256 characters
// - must contain at least one uppercase letter
// - must contain at least one lowercase letter
// - must contain at least one number
// - must contain at least one special character
var (
	passwordRegex     = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+-=[]{}|;:,.<>?]+$`)
	passwordMinLength = 8
	passwordMaxLength = 256
)

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < passwordMinLength || len(password) > passwordMaxLength {
		return false
	}

	if !passwordRegex.MatchString(password) {
		return false
	}
	return true
}

// validateEmail checks if the email is valid
// ruleset:
// - must be a valid email address
var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	return emailRegex.MatchString(email)
}
