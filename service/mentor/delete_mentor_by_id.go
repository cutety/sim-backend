package mentor

import "sim-backend/models"

type DeleteMentorByIDService struct {}

// DeleteMentorByID 删除mentor
func (*DeleteMentorByIDService) DeleteMentorByID(userID string) error {
	// Do something...
	return models.MMentor.DeleteMentorByID(userID)
}