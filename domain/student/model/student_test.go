package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudentWithNotExistEmail(t *testing.T) {
	student := NewStudent()
	student.ID = 1
	student.Name = "Aluno Teste"
	student.Password = "1234"
	err := student.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Email is required", err.Error())
}

func TestStudentWithNotExistPassword(t *testing.T) {
	student := NewStudent()
	student.ID = 1
	student.Name = "Aluno Teste"
	student.Email = "student@test.com"
	err := student.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Password is required", err.Error())
}
