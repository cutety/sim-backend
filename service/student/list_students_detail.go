package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
)

type ListStudentsDetailService struct {

}

func (*ListStudentsDetailService) ListStudentsDetail(pagination *common.Pagination) ([]models.StudentDetail, int64, error) {
	apps, total, err := models.MStudent.GetDetailByStuID(pagination)
	return apps, total, err
}

