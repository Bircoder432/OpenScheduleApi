package dto

import (
	"strings"
	"time"
)

type CollegeResponse struct {
	ID       uint             `json:"collegeId"`
	Name     string           `json:"name"`
	Campuses []CampusResponse `json:"campuses,omitempty"`
}

type CampusResponse struct {
	ID            uint                   `json:"campusId"`
	Name          string                 `json:"name"`
	StudentGroups []StudentGroupResponse `json:"groups,omitempty"`
}
type StudentGroupResponse struct {
	ID       uint   `json:"studentGroupId"`
	Name     string `json:"name"`
	CampusID uint   `json:"campusId"`
}

type Lesson struct {
	LessonID       uint      `json:"lessonId"`
	Title          string    `json:"title"`
	Cabinet        string    `json:"cabinet"`
	Date           time.Time `json:"date"`
	Teacher        string    `json:"teacher"`
	Order          uint      `json:"order"`
	StudentGroupID uint      `json:"studentGroupID"`
}

type Call struct {
	CallID  uint         `json:"callId"`
	Weekday time.Weekday `json:"weekday"`
	Begins  HourMinute   `json:"begins"`
	Ends    HourMinute   `json:"ends"`
	Order   uint         `json:"order"`
}

type HourMinute time.Time

func (hm *HourMinute) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("15:04", s)
	if err != nil {
		return err
	}
	*hm = HourMinute(t)
	return nil
}

func (hm HourMinute) MarshalJSON() ([]byte, error) {
	t := time.Time(hm)
	s := t.Format("15:04:05")
	return []byte(`"` + s + `"`), nil
}
