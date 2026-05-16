package helpers

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateFullName(name string) error {
	name = strings.TrimSpace(name)
	if len(name) < 2 {
		return errors.New("full name must be at least 2 characters")
	}
	if len(name) > 100 {
		return errors.New("full name must not exceed 100 characters")
	}
	if matched, _ := regexp.MatchString(`^[\p{L}\s'-]+$`, name); !matched {
		return errors.New("full name contains invalid characters")
	}
	return nil
}

func ValidateUsername(username string) error {
	username = strings.TrimSpace(username)
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if len(username) > 30 {
		return errors.New("username must not exceed 30 characters")
	}
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, username); !matched {
		return errors.New("username can only contain letters, numbers, underscore and hyphen")
	}
	return nil
}