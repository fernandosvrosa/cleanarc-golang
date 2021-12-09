package config

import (
	"database/sql"
	"github.com/fernandosvrosa/cleanarc-golang/application/student/entrypoint"
	"github.com/fernandosvrosa/cleanarc-golang/application/student/gateway"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/usecase"
)

type Factory struct {
	db *sql.DB
}

func NewFactory(db *sql.DB) *Factory {
	return &Factory{db: db}
}

func (s *Factory) CreateStudentHandler() *entrypoint.StudentHandler {
	provider := gateway.NewInsertStudentGatawayProvider(s.db)
	saveStudentUseCase := usecase.NewSaveStudent(provider)
	return entrypoint.NewStudentHandler(saveStudentUseCase)
}
