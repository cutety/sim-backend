package course

import "sim-backend/models"

type ListCoursesService struct {

}

func (*ListCoursesService) ListService(mentorID, grade, class string) ([]models.Course, error) {
	return models.MCourse.ListCourses(mentorID, grade, class)
}
