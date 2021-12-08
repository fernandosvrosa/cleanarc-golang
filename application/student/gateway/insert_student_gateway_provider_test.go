package gateway

import (
	"github.com/fernandosvrosa/cleanarc-golang/application/student/gateway/fixture"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSaveStudentSuccess(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)

	input := model.Student{
		Name:     "Test",
		Email:    "student@test.com",
		Password: "202cb962ac59075b964b07152d234b70",
	}

	provider := NewInsertStudentGatawayProvider(db)
	result, err := provider.InsertStudent(input)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
}
