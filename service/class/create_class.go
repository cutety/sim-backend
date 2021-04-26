package class

import (
	"sim-backend/models"
	"sim-backend/utils"
)

type CreateClassService struct {
	Grade string `json:"grade" validate:"required" label:"年级"`
	Class string `json:"class" validate:"required" label:"班级"`
}

func (s *CreateClassService) CreateClass() error {
	classID := utils.UUID()
	class := models.Class{
		ClassID: classID,
		Grade: s.Grade,
		Name: s.Class,
	}
	return class.Create()
}
