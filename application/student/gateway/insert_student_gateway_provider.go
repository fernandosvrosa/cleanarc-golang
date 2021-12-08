package gateway

import (
	"database/sql"
	"github.com/fernandosvrosa/cleanarc-golang/application/student/gateway/repository"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/model"
)

type InsertStudentGatawayProvider struct {
	db *sql.DB
}

func NewInsertStudentGatawayProvider(db *sql.DB) *InsertStudentGatawayProvider {
	return &InsertStudentGatawayProvider{db: db}
}

func (i *InsertStudentGatawayProvider) InsertStudent(student model.Student) (model.Student, error) {

	repository := repository.NewStudentRepository(i.db)

	if err := repository.Insert(student); err != nil {
		return model.Student{}, err
	}

	result, err := repository.FindByEmail(student.Email)

	if err != nil {
		return model.Student{}, err
	}

	return result, nil
}
