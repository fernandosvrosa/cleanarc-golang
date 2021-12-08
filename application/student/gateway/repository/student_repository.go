package repository

import (
	"database/sql"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/model"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (s *StudentRepository) Insert(student model.Student) error {
	stmt, err := s.db.Prepare(`
		INSERT INTO student(name,email,password) 
		values ($1, $2, $3)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(student.Name, student.Email, student.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *StudentRepository) FindByEmail(email string) (model.Student, error) {
	row, err := s.db.Query(`
		select id, name, email, password from student
		where email = $1
	`, email)
	if err != nil {
		return model.Student{}, err
	}
	defer row.Close()

	var student = model.Student{}
	for row.Next() { // Iterate and fetch the records from result cursor
		row.Scan(&student.ID, &student.Name, &student.Email, &student.Password)
	}
	return student, nil
}
