package usecase

import (
	"crypto/md5"
	"encoding/hex"
	mock_gateway "github.com/fernandosvrosa/cleanarc-golang/domain/student/gateway"
	"testing"

	"github.com/fernandosvrosa/cleanarc-golang/domain/student/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSaveStudent(t *testing.T) {
	password := []byte("123")
	hashedPassword := md5.Sum(password)
	passwordString := hex.EncodeToString(hashedPassword[:])
	input := model.Student{
		Name:     "Test",
		Email:    "student@test.com",
		Password: "123",
	}

	inputMock := model.Student{
		Name:     "Test",
		Email:    "student@test.com",
		Password: passwordString,
	}

	expectedOutput := model.Student{
		ID:       1,
		Name:     "Test",
		Email:    "student@test.com",
		Password: passwordString,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := mock_gateway.NewMockInsertStudentGateway(ctrl)
	repositoryMock.EXPECT().InsertStudent(inputMock).Return(expectedOutput, nil)

	usecase := NewSaveStudent(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestSaveStudentInvalidEmail(t *testing.T) {
	input := model.Student{
		Name:     "Test",
		Email:    "",
		Password: "1234",
	}

	usecase := NewSaveStudent(nil)
	_, err := usecase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, "Email is required", err.Error())

}
