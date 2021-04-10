package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"sim-backend/models/common"
	"time"
)

var MStudent Student

type Student struct {
	ID             uint       `gorm:"primary_key"`
	StuID          string     `gorm:"column:stu_id;type:varchar(20)" json:"stu_id" validate:"required" label:"用户ID"`
	StuName        string     `gorm:"column:stu_name;type:varchar(20)" json:"stu_name" validate:"required" label:"姓名"`
	Gender         string     `gorm:"column:gender;type:int(1);not null;default:1" json:"gender"`
	Birthday       time.Time  `gorm:"column:birthday;type:timestamp" json:"birthday"`
	PolicalStatus  string     `gorm:"column:polical_status;type:varchar(45)" json:"polical_status"`
	Nation         string     `gorm:"column:nation;type:varchar(4)" json:"nation"`
	Grade          string     `gorm:"column:grade;type:varchar(4)" json:"grade"`
	AdmissionMajor string     `gorm:"column:admission_major;type:varchar(50)" json:"admission_major"`
	Phone          string     `gorm:"column:phone;type:varchar(20);" json:"phone" validate:"required" label:"电话号"`
	Email          string     `gorm:"column:email;type:varchar(50);" json:"email"`
	Wechat         string     `grom:"column:wechat;type:varchar(25);" json:"wechat"`
	QQ             string     `gorm:"column:qq;type:varchar(10);" json:"qq"`
	CreatedAt      time.Time  `gorm:"type:timestamp"`
	UpdatedAt      time.Time  `gorm:"type:timestamp"`
	DeletedAt      *time.Time `sql:"index" gorm:"type:timestamp"`
}

func (*Student) TableName() string {
	return "students"
}

func (*Student) Update(student *Student) error {
	return extension.DB.Model(&student).Where("stu_id = ?", student.StuID).Updates(&student).Error
}

func (s *Student) GetStudentsAmountByGrade(grade string) (int64, error) {
	var total int64
	err := extension.DB.Table(s.TableName()).Where("grade = ?", grade).Count(&total).Error
	return total, err
}

type StudentsInfoTable struct {
	StuName        string
	StuID          string
	AdmissionMajor string `json:"major"`
}

func (c *Student) GetStudentsInfoTable(grade string) ([]StudentsInfoTable, error) {
	var apps []StudentsInfoTable
	err := extension.DB.Table(c.TableName()).Where("grade = ?", grade).Find(&apps).Error
	return apps, err
}

func (*Student) GetByStuID(stuID string) (*Student, error) {
	stu := &Student{}
	err := extension.DB.Where("stu_id = ?", stuID).Find(&stu).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return stu, err
}

type StudentDetail struct {
	UserID            string  `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	StuName           string  `gorm:"column:stu_name;type:varchar(20)" json:"stu_name" validate:"required" label:"姓名"`
	Gender            string  `gorm:"column:gender;type:int(1);not null;default:1" json:"gender"`
	Grade             string  `gorm:"column:grade;type:varchar(4)" json:"grade"`
	Major             string  `gorm:"column:major;type:varchar(50)" json:"major"`
	Phone             string  `gorm:"column:phone;type:varchar(20);" json:"phone" validate:"required" label:"电话号"`
	Email             string  `gorm:"column:email;type:varchar(50);" json:"email"`
	Wechat            string  `grom:"column:wechat;type:varchar(255);" json:"wechat"`
	QQ                string  `gorm:"column:qq;type:varchar(10);" json:"qq"`
	MentorUserID      string  `gorm:"column:mentor_user_id;type:varchar(20)" json:"mentor_user_id" label:"导师用户ID"`
	MentorName        string  `gorm:"column:mentor_name;type:varchar(20);" json:"mentor_name" label:"导师姓名"`
	ApplySchool       string  `gorm:"column:apply_school;type:varchar(255)" json:"apply_school" validate:"required" label:"报考院校"`
	ApplyMajor        string  `gorm:"column:apply_major;type:varchar(255)" json:"apply_major" validate:"required" label:"报考专业"`
	PreliminaryResult float32 `gorm:"column:preliminary_result;type:decimal(11,2)" json:"preliminary_result" label:"初试成绩"`
	RetrailResult     float32 `gorm:"column:retrail_result;type:decimal(11,2)" json:"retrail_result" label:"复试成绩"`
	AdmissionSchool   string  `gorm:"column:admission_school;type:varchar(255)" json:"admission_school" validate:"required" label:"录取院校"`
	AdmissionMajor    string  `gorm:"column:admission_major;type:varchar(255)" json:"admission_major" validate:"required" label:"录取院校"`
	IsAdmitted        bool    `gorm:"column:is_admitted;type:tinyint(1)" json:"is_admitted" label:"录取结果"`
}

func (s *Student) GetDetailByStuID(stuName string, gender int, grade string, major string, pagination *common.Pagination) ([]StudentDetail, int64, error) {
	var total int64
	var apps []StudentDetail
	sql := `
			SELECT 
				s.stu_id as user_id, s.stu_name as stu_name, s.gender as gender, s.grade as grade, s.admission_major as major,
				s.qq as qq, s.phone as phone, s.email as email, s.wechat as wechat, a.mentor_user_id as mentor_user_id, m.name as mentor_name,
				a.apply_school as apply_school, a.apply_major as apply_major, a.preliminary_result as preliminary_result,
				a.retrail_result as retrail_result, a.admission_school as admission_school, a.admission_major as admission_major,
				a.is_admitted as is_admitted
			FROM
				students s
			LEFT JOIN
				application a
					ON a.user_id = s.stu_id
			LEFT JOIN
				mentors m
					ON m.user_id = a.mentor_user_id
`
	var whereParams []interface{}
	whereSQL := `
		WHERE	1 = 1
`
	if stuName != "" {
		whereSQL += `
			AND s.stu_name like ?
`
		whereParams = append(whereParams, stuName+"%")
	}
	if gender != -1 {
		whereSQL += `
			AND s.gender = ?
`
		whereParams = append(whereParams, gender)
	}
	if grade != "" {
		whereSQL += `
			AND s.grade = ?
`
		whereParams = append(whereParams, grade)
	}
	if major != "" {
		whereSQL += `
			AND s.admission_major = ?
`
		whereParams = append(whereParams, major)
	}
	total = extension.DB.Raw(sql+whereSQL, whereParams...).Scan(&apps).RowsAffected
	whereParams = append(whereParams, pagination.Limit)
	whereParams = append(whereParams, (pagination.Page-1)*pagination.Limit)
	err := extension.DB.Raw(sql+whereSQL+`LIMIT ? OFFSET ?`, whereParams...).
		Scan(&apps).Error
	if err != nil {
		return nil, 0, err
	}
	return apps, total, err
}

type MaleAndFemaleAmount struct {
	Gender string
	Amount int64
}

// GetAmountByGender 获取男女人数
func (s *Student) GetAmountByGender(grade string, gender int) (int64, error) {
	var total int64
	err := extension.DB.Table(s.TableName()).Where("grade = ? and gender = ?", grade, gender).Count(&total).Error
	return total, err
}

type StudentsValue struct {
	Name  string `json:"name"`
	Value int64 `json:"value"`
}


// GetAgeDistribution 获得年龄分布
func (s *Student) GetAgeDistribution(grade string) ([]StudentsValue, error) {
	var result []StudentsValue
	err := extension.DB.Raw(`
		SELECT 
			TIMESTAMPDIFF(YEAR, birthday, CURDATE()) AS name,
			COUNT(*) value
		FROM 
			students
		WHERE 
			birthday IS NOT NULL	
			AND grade = ?
		GROUP BY name
		ORDER BY name ASC
	`, grade).Scan(&result).Error
	return result, err
}


// GetStudentsProvince 获取省份
func (s *Student) GetStudentsProvince(grade string) ([]StudentsValue, error) {
	var results []StudentsValue
	sql := `
	SELECT CASE
    WHEN adcode >= '110000' AND adcode < '120000' THEN '北京'
    WHEN adcode >= '120000' AND adcode < '130000' THEN '天津'
    WHEN adcode >= '130000' AND adcode < '140000' THEN '河北'
    WHEN adcode >= '140000' AND adcode < '150000' THEN '山西'
    WHEN adcode >= '150000' AND adcode < '210000' THEN '内蒙古'
    WHEN adcode >= '210000' AND adcode < '220000' THEN '辽宁'
    WHEN adcode >= '220000' AND adcode < '230000' THEN '吉林'
    WHEN adcode >= '230000' AND adcode < '310000' THEN '黑龙江'
    WHEN adcode >= '310000' AND adcode < '320000' THEN '上海'
    WHEN adcode >= '320000' AND adcode < '330000' THEN '江苏'
    WHEN adcode >= '330000' AND adcode < '340000' THEN '浙江'
    WHEN adcode >= '340000' AND adcode < '350000' THEN '安徽'
    WHEN adcode >= '350000' AND adcode < '360000' THEN '福建'
    WHEN adcode >= '360000' AND adcode < '370000' THEN '江西'
    WHEN adcode >= '370000' AND adcode < '410000' THEN '山东'
    WHEN adcode >= '410000' AND adcode < '420000' THEN '河南'
    WHEN adcode >= '420000' AND adcode < '430000' THEN '湖北'
    WHEN adcode >= '430000' AND adcode < '440000' THEN '湖南'
    WHEN adcode >= '440000' AND adcode < '450000' THEN '广东'
    WHEN adcode >= '450000' AND adcode < '460000' THEN '广西'
    WHEN adcode >= '460000' AND adcode < '500000' THEN '海南'
    WHEN adcode >= '500000' AND adcode < '510000' THEN '重庆'
    WHEN adcode >= '510000' AND adcode < '520000' THEN '四川'
    WHEN adcode >= '520000' AND adcode < '530000' THEN '贵州'
    WHEN adcode >= '530000' AND adcode < '540000' THEN '云南'
    WHEN adcode >= '540000' AND adcode < '610000' THEN '西藏'
    WHEN adcode >= '610000' AND adcode < '620000' THEN '陕西'
    WHEN adcode >= '620000' AND adcode < '630000' THEN '甘肃'
    WHEN adcode >= '630000' AND adcode < '640000' THEN '青海'
    WHEN adcode >= '640000' AND adcode < '650000' THEN '宁夏'
    WHEN adcode >= '650000' AND adcode < '710000' THEN '新疆'
    WHEN adcode >= '710000' AND adcode < '810000' THEN '台湾'
    WHEN adcode >= '810000' AND adcode < '820000' THEN '香港'
    WHEN adcode >= '820000' THEN '澳门'
    END AS name,
    COUNT(*) AS value
    FROM students
    where grade = ?
    GROUP BY name
`
	err := extension.DB.Raw(sql, grade).Scan(&results).Error
	return results, err
}

// GetFirstnameByGrade 根据年级获取姓排行
func (s *Student) GetFirstnameByGrade(grade string, count int) ([]StudentsValue, error) {
	var result []StudentsValue
	sql := `
		SELECT
			COUNT(*) value,
			name
				FROM (
					SELECT 
						LEFT(stu_name, 1) name
					FROM 
						students
					WHERE grade = ?
					) t
		GROUP BY NAME
		ORDER BY VALUE DESC
	`
	var whereVals []interface{}
	whereVals = append(whereVals, grade)
	if count > 0 {
		sql += `limit ?`
		whereVals = append(whereVals, count)
	}
	err := extension.DB.Raw(sql, whereVals...).Scan(&result).Error
	return result, err
}

// GetSameNameByGrade 根据年纪获取同名情况
func (s *Student) GetSameNameByGrade(grade string) ([]StudentsValue, error) {
	var result []StudentsValue
	err := extension.DB.Table(s.TableName()).
		Select("count(*) value, stu_name as name").
		Where("grade = ? ", grade).
		Group("name").
		Order("value desc").
		Find(&result).Error

	return result, err
}

func (*Student) GetSameBirthdayByGrade(grade string)([]StudentsValue, error) {
	var result []StudentsValue
	sql := `
		SELECT 
			COUNT(*) value,
			SUBSTRING(birthday,6, 5) name
		FROM 
			students
		WHERE 
			birthday IS NOT NULL 
			AND grade = ?
		GROUP BY name
		HAVING VALUE > 3
		ORDER BY name ASC 
		`
	err := extension.DB.Raw(sql, grade).Scan(&result).Error
	return result, err
}

// GetMajorRankByGrade 根据年级获取专业排行
func (*Student) GetMajorRankByGrade(grade string)([]StudentsValue, error) {
	var result []StudentsValue
	sql := `
	SELECT 
		admission_major AS name,
		COUNT(admission_major) AS value
	FROM 
		students
	WHERE
		grade = ?
		AND admission_major != ""
	GROUP BY 
		admission_major
	ORDER BY 
		value DESC
		`
	err := extension.DB.Raw(sql, grade).Scan(&result).Error
	return result, err
}

