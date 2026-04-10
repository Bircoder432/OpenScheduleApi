package dto

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type NewParserRequest struct {
	CollegeName string   `json:"collegeName"`
	CampusNames []string `json:"campusNames"`
}

type NewParserResponse struct {
	Token string `json:"token"`
}

type UpdateGroupsRequest struct {
	CampusID          uint     `json:"campusId"`
	StudentGroupNames []string `json:"studentGroupNames"`
}

type GetParserResponse struct {
	CollegeID uint `json:"collegeId"`
}

type ScheduleResponse struct {
	GroupID uint                     `json:"groupId"`
	Date    time.Time                `json:"date"`
	Lessons []ScheduleLessonResponse `json:"lessons"`
}

type ScheduleLessonResponse struct {
	Title     string     `json:"title"`
	Cabinet   string     `json:"cabinet"`
	Teacher   string     `json:"teacher"`
	Order     uint       `json:"order"`
	StartTime HourMinute `json:"startTime"`
	EndTime   HourMinute `json:"endTime"`
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
