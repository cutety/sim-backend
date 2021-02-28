package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"sim-backend/models/common"
	"time"
)

var MUser User

type User struct {
	ID        uint `gorm:"primary_key"`
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id"`
	Username string `gorm:"column:username;type:varchar(20)" json:"username"`
	Password string `gorm:"column:password;type:varchar(20)" json:"password"`
	Role int `gorm:"column:role;type:int;DEFAULT:2" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt *time.Time `sql:"index" gorm:"type:timestamp"`
}

func (*User) TableName() string {
	return "users"
}

func (*User) GetUserByUserID(userID string) (*User, error) {
	user := &User{}
	err := extension.DB.Where("user_id = ?", userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &User{}, err
	}
	if err != nil {
		return nil, err
	}
	return user, err
}

func (*User) GetUserByID(id uint) (*User, error) {
	user := &User{}
	err := extension.DB.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &User{}, nil
	}
	return user, err
}

func (*User) UpdatePasswordById(id uint, password string) error {
	return extension.DB.Table(MUser.TableName()).Where("id = ?", id).Update("password", password).Error
}

func (*User) Total() (*int, error) {
	var total *int
	err := extension.DB.Table(MUser.TableName()).Count(&total).Error
	return total, err
}

type MatchingResult struct {
	MatchSchool string `json:"match_school"`
	MatchDegree string `json:"match_degree"`
	MatchMajor string `json:"match_major"`
	Status string `json:"status"`
	Mentor
}

func (*User) GetMatchingResult(pagination *common.Pagination, userID string) ([]MatchingResult, int64, error) {
	var result []MatchingResult
	var total int64
	err := extension.DB.Raw(`
			SELECT
			a.apply_school,
			case 
				when a.apply_school = m.undergraduate_university then '学士'
				when a.apply_school = m.graduate_school then '硕士'
				when a.apply_school = m.phd_school then '博士'
				when a.apply_major = m.undergraduate_major then '学士'
				when a.apply_major = m.graduate_major then '硕士'
				when a.apply_major = m.phd_major then '博士'
			end as 'match_degree',
			case 
				when a.apply_school = m.undergraduate_university then m.undergraduate_major
				when a.apply_school = m.graduate_school then m.graduate_major
				when a.apply_school = m.phd_school then m.phd_major
				when a.apply_major = m.undergraduate_major then m.undergraduate_major
				when a.apply_major = m.graduate_major then m.graduate_major
				when a.apply_major = m.phd_major then m.phd_major
			end as 'match_major',
			case a.apply_major
				when m.undergraduate_major then m.undergraduate_university 
				when m.graduate_major then m.graduate_school
				when m.phd_major then m.phd_school
			end as 'match_school',
			m.user_id ,
			case
			when m.user_id = a.mentor_user_id then 1
			else 0
			end as 'status',
			m.*
		FROM
			mentors m
			left join
			application a
			on a.apply_school = m.undergraduate_university
			or a.apply_school = m.graduate_school
			or a.apply_school = m.phd_school
			or a.apply_major = m.undergraduate_major 
			or a.apply_major = m.graduate_major 
			or a.apply_major = m.phd_major 
			left join
			users u
			on u.user_id = a.user_id
		WHERE u.user_id = ?
		LIMIT ? OFFSET ?
    `, userID, pagination.Limit, (pagination.Page - 1) * pagination.Limit).
		Scan(&result).Error
	total = int64(len(result))
	return result, total, err
}