package dto

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

type UpdateGroupsRequest struct {
	CampusID          uint     `json:"campusId"`
	StudentGroupNames []string `json:"studentGroupNames"`
}

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

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}

func NewErrorResponse(err string, statusCode int) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		Error:      err,
	}
}

func (e ErrorResponse) Send(ctx fiber.Ctx) error {
	return ctx.Status(e.StatusCode).JSON(e)
}
