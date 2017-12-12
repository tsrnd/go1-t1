package controllers

import (
	"fmt"
	"goweb1/messages"
	"regexp"
)

type Validator struct {
	err string
}

func Character(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(s)
}

func email(s string) bool {
	return regexp.MustCompile(".+@.+\\..+").Match([]byte(s))
}

func CharacterAll(s string) bool {
	return regexp.MustCompile(`^[/s]+$`).MatchString(s)
}

func LenString(min, max int, s string) bool {
	return !(len(s) < min || len(s) > max)
}

func (v *Validator) ValidatePassword(s string) bool {
	if !LenString(6, 255, s) {
		v.err = "Password " + fmt.Sprintf(messages.Validate_length_string, 6, 255)
		return false
	}
	return true
}

func (v *Validator) ValidateUsername(s string) bool {
	if !LenString(6, 255, s) {
		v.err = "Username " + fmt.Sprintf(messages.Validate_length_string, 6, 255)
		return false
	}
	if !Character(s) {
		v.err = "Username " + messages.Validate_character
		return false
	}

	return true
}

func (v *Validator) ValidateFullName(s string) bool {
	if !LenString(6, 255, s) {
		v.err = "Fullname " + fmt.Sprintf(messages.Validate_length_string, 6, 255)
		return false
	}
	if !CharacterAll(s) {
		v.err = "Fullname " + messages.Validate_character
		return false
	}

	return true
}

func (v *Validator) ValidateAddress(s string) bool {
	if !LenString(6, 255, s) {
		v.err = "Address " + fmt.Sprintf(messages.Validate_length_string, 6, 255)
		return false
	}
	if !CharacterAll(s) {
		v.err = "Address " + messages.Validate_character
		return false
	}

	return true
}

func (v *Validator) ValidateEmail(s string) bool {
	if !LenString(6, 255, s) {
		v.err = "Email " + fmt.Sprintf(messages.Validate_length_string, 6, 255)
		return false
	}
	if !email(s) {
		v.err = "Email " + messages.Validate_email
		return false
	}

	return true
}

func (v *Validator) ExistsEmail(email string) bool {
	// user, _ := model.GetUserByEmail(email)
	// if (user.Email != "") {
	// 	v.err = "Email " + fmt.Sprintf(messages.Validate_exists)
	// 	return true
	// }
	return false
}

func (v *Validator) ExistsUserName(username string) bool {
	// user, _ := model.GetUserByUserName(username)
	// if (user.Username != "") {
	// 	v.err = "Username " + fmt.Sprintf(messages.Validate_exists)
	// 	return true
	// }
	return false
}
