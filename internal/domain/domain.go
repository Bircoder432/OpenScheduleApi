package domain

import (
	"time"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/dto"
)

type Parser struct {
	ID        uint
	Token     string
	CollegeID uint
}

type College struct {
	ID    uint
	Name  string
	Token string
}

type Campus struct {
	ID        uint
	Name      string
	CollegeID uint
}

type StudentGroup struct {
	ID       uint
	Name     string
	CampusID uint
}

type Call struct {
	ID        uint
	Weekday   time.Weekday
	Begins    time.Time
	Ends      time.Time
	Order     uint
	CollegeID uint
}

type Lesson struct {
	ID             uint
	Title          string
	Cabinet        string
	Date           time.Time
	Teacher        string
	Order          uint
	StudentGroupID uint
}

func (c College) ToDTO(campuses []Campus) dto.CollegeResponse {
	campusResponses := make([]dto.CampusResponse, 0, len(campuses))
	for _, campus := range campuses {
		campusResponses = append(campusResponses, campus.ToDTO(nil))
	}
	return dto.CollegeResponse{
		ID:       c.ID,
		Name:     c.Name,
		Campuses: campusResponses,
	}
}

func (c Campus) ToDTO(studentGroups []StudentGroup) dto.CampusResponse {
	studentGroupResponses := make([]dto.StudentGroupResponse, 0, len(studentGroups))
	for _, studentGroup := range studentGroups {
		studentGroupResponses = append(studentGroupResponses, studentGroup.ToDTO())
	}
	return dto.CampusResponse{
		ID:            c.ID,
		Name:          c.Name,
		StudentGroups: studentGroupResponses,
	}
}

func (s StudentGroup) ToDTO() dto.StudentGroupResponse {
	return dto.StudentGroupResponse{ID: s.ID, Name: s.Name, CampusID: s.CampusID}
}
