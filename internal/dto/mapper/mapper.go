package mapper

import (
	"time"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/dto"
)

func NewSchedule(lessons []domain.Lesson, calls []domain.Call) dto.ScheduleResponse {
	if len(lessons) == 0 {
		return dto.ScheduleResponse{}
	}

	mapCalls := make(map[time.Weekday]map[uint]domain.Call)

	for _, call := range calls {
		if mapCalls[call.Weekday] == nil {
			mapCalls[call.Weekday] = make(map[uint]domain.Call)
		}
		mapCalls[call.Weekday][call.Order] = call
	}
	lessonResponses := make([]dto.ScheduleLessonResponse, 0, len(lessons))
	for _, lesson := range lessons {
		lessonResponses = append(lessonResponses, newScheduleLesson(lesson, mapCalls))
	}
	scheduleResponse := dto.ScheduleResponse{
		GroupID: lessons[0].StudentGroupID,
		Date:    lessons[0].Date,
		Lessons: lessonResponses,
	}
	return scheduleResponse
}

func NewSchedules(lessons []domain.Lesson, calls []domain.Call) []dto.ScheduleResponse {
	if len(lessons) == 0 {
		return nil
	}

	mapCalls := make(map[time.Weekday]map[uint]domain.Call)

	for _, call := range calls {
		if mapCalls[call.Weekday] == nil {
			mapCalls[call.Weekday] = make(map[uint]domain.Call)
		}
		mapCalls[call.Weekday][call.Order] = call
	}
	lessonResponses := make(map[time.Time][]dto.ScheduleLessonResponse)
	for _, lesson := range lessons {
		lessonResponses[lesson.Date] =
			append(lessonResponses[lesson.Date], newScheduleLesson(lesson, mapCalls))
	}
	scheduleResponses := []dto.ScheduleResponse{}
	for t, resps := range lessonResponses {
		scheduleResponses = append(scheduleResponses, dto.ScheduleResponse{
			GroupID: lessons[0].StudentGroupID, Date: t,
			Lessons: resps,
		})
	}
	return scheduleResponses
}

func newScheduleLesson(lesson domain.Lesson, mapCalls map[time.Weekday]map[uint]domain.Call) dto.ScheduleLessonResponse {
	return dto.ScheduleLessonResponse{
		Title:     lesson.Title,
		Cabinet:   lesson.Cabinet,
		Teacher:   lesson.Teacher,
		Order:     lesson.Order,
		StartTime: dto.HourMinute(mapCalls[lesson.Date.Weekday()][lesson.Order].Begins),
		EndTime:   dto.HourMinute(mapCalls[lesson.Date.Weekday()][lesson.Order].Ends),
	}
}
