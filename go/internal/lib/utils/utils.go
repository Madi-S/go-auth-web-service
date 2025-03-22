package utils

import (
	"errors"
	"net"
	"net/mail"
	"strings"
)

func ValidateEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return errors.New("invalid email")
	}

	if _, err := net.LookupMX(parts[1]); err != nil {
		return err
	}

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("empty password")
	}
	if len(password) < 5 || len(password) > 60 {
		return errors.New("password should be between 5 and 60 chars long")
	}

	return nil
}

func ValidateAppId(appId int32) error {
	if appId == 0 {
		return errors.New("app id is empty")
	}

	return nil
}
