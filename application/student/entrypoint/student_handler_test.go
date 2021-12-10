package entrypoint

import (
	"bytes"
	"encoding/json"
	"github.com/fernandosvrosa/cleanarc-golang/application/student/gateway"
	"github.com/fernandosvrosa/cleanarc-golang/config/fixture"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestStudentHandler(t *testing.T) {
	migrationDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationDir)
	defer fixture.Down(db, migrationDir)

	provider := gateway.NewInsertStudentGatawayProvider(db)
	saveStudentUseCase := usecase.NewSaveStudent(provider)
	studentHandler := NewStudentHandler(saveStudentUseCase)

	app := fiber.New()
	app.Post("/v1/students", studentHandler.SaveStudent)

	body := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{"Teste", "teste@student.comm", "1234"}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(body)
	req := httptest.NewRequest(http.MethodPost, "/v1/students", &b)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, 600)

	assert.Equalf(t, 201, resp.StatusCode, "Save student")
}
