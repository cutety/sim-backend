package user

import "sim-backend/models"

type ListAllMentorService struct {

}

func (*ListAllMentorService) ListAllMentor() ([]models.Mentor, int64, error) {
	mentors, total, err := models.MMentor.ListAllMentors()
	if err != nil {
		return nil, 0, err
	}
	return mentors, total, err
}
