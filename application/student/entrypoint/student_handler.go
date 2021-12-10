package entrypoint

import (
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/model"
	"github.com/fernandosvrosa/cleanarc-golang/domain/student/usecase"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	saveStudent *usecase.SaveStudent
}

func NewStudentHandler(saveStudent *usecase.SaveStudent) *StudentHandler {
	return &StudentHandler{saveStudent: saveStudent}
}

type StudentOutDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *StudentHandler) SaveStudent(c *fiber.Ctx) error {
	student := &model.Student{}
	if err := c.BodyParser(student); err != nil {
		return err
	}
	result, err := s.saveStudent.Execute(*student)
	if err != nil {
		return err
	}
	studentOut := StudentOutDTO{result.ID, result.Name, result.Email}

	return c.Status(fiber.StatusCreated).JSON(studentOut)
}
