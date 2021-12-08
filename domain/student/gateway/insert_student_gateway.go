package gateway

import "github.com/fernandosvrosa/cleanarc-golang/domain/student/model"

type InsertStudentGateway interface {
	InsertStudent(student model.Student) (model.Student, error)
}
