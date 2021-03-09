package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
)

type ListStudentsDetailService struct {
	StuName string `form:"stu_name"`
	Gender int `form:"gender"`
	Grade string `form:"grade"`
	Major string `form:"major"`
}

func (s *ListStudentsDetailService) ListStudentsDetail(pagination *common.Pagination) ([]models.StudentDetail, int64, error) {
	apps, total, err := models.MStudent.GetDetailByStuID(s.StuName, s.Gender, s.Grade, s.Major ,pagination)
	return apps, total, err
}

