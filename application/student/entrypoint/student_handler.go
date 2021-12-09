package entrypoint

import (
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/usecase"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	saveStudent *usecase.SaveStudent
}

func NewStudentHandler(saveStudent *usecase.SaveStudent) *StudentHandler {
	return &StudentHandler{saveStudent: saveStudent}
}

func (s *StudentHandler) SaveStudent(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"examIds":  1,
		"accounts": 1,
	})
}
