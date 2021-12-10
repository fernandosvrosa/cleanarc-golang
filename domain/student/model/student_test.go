package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStudentWithNotExistEmail(t *testing.T) {
	student := NewStudent(1, "Aluno Teste", "", "1234")
	err := student.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Email is required", err.Error())
}

func TestStudentWithNotExistPassword(t *testing.T) {
	student := NewStudent(1, "Aluno Teste", "student@test.com", "")
	err := student.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Password is required", err.Error())
}
