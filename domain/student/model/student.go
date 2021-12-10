package model

import (
	"errors"
	"strings"
)

type Student struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func NewStudent(id int, name, email, password string) *Student {
	return &Student{ID: id, Name: name, Email: email, Password: password}
}

func (s *Student) IsValid() error {
	if "" == strings.TrimSpace(s.Email) {
		return errors.New("Email is required")
	}

	if "" == strings.TrimSpace(s.Password) {
		return errors.New("Password is required")
	}

	return nil
}
