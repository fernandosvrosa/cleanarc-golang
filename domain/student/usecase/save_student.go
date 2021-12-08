package usecase

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/gateway"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/model"
)

type SaveStudent struct {
	gateway.InsertStudentGateway
}

func NewSaveStudent(insertStudentGateway gateway.InsertStudentGateway) *SaveStudent {
	return &SaveStudent{InsertStudentGateway: insertStudentGateway}
}

func (saveStudent *SaveStudent) Execute(student model.Student) (model.Student, error) {
	password := []byte(student.Password)
	hashedPassword := md5.Sum(password)
	student.Password = hex.EncodeToString(hashedPassword[:])

	err := student.IsValid()
	if err != nil {
		return model.Student{}, err
	}
	return saveStudent.InsertStudentGateway.InsertStudent(student)
}
