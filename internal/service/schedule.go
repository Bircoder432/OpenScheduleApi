package service

import (
	"context"
	"time"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/dto"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/dto/mapper"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/weekdays"
)

type ScheduleService struct {
	lessonRepo  repository.LessonRepo
	callRepo    repository.CallRepo
	groupRepo   repository.StudentGroupRepo
	collegeRepo repository.CollegeRepo
}

func NewScheduleService(
	groupRepo repository.StudentGroupRepo,
	lessonRepo repository.LessonRepo,
	callRepo repository.CallRepo,
	collegeRepo repository.CollegeRepo) *ScheduleService {
	return &ScheduleService{
		lessonRepo:  lessonRepo,
		groupRepo:   groupRepo,
		callRepo:    callRepo,
		collegeRepo: collegeRepo,
	}
}

func (s *ScheduleService) GetScheduleByDay(ctx context.Context, groupID uint, day string) (dto.ScheduleResponse, error) {
	switch day {
	case "today":
		return s.GetScheduleForDate(ctx, groupID, time.Now())
	case "tomorrow":
		return s.GetScheduleForDate(ctx, groupID, time.Now().AddDate(0, 0, 1))
	default:
		return dto.ScheduleResponse{}, domain.ErrInvalidData
	}
}

func (s *ScheduleService) GetScheduleForDate(ctx context.Context, groupID uint, date time.Time) (dto.ScheduleResponse, error) {
	lessons, err := s.lessonRepo.GetForDate(ctx, groupID, date)
	if err != nil {
		return dto.ScheduleResponse{}, err
	}
	calls, err := s.getCallsByGroup(ctx, groupID)
	if err != nil {
		return dto.ScheduleResponse{}, err
	}
	schedule := mapper.NewSchedule(lessons, calls)
	return schedule, nil
}

func (s *ScheduleService) GetSchedulesForDates(ctx context.Context, groupID uint, startDate, endDate time.Time) ([]dto.ScheduleResponse, error) {
	lessons, err := s.lessonRepo.GetForDates(ctx, groupID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	calls, err := s.getCallsByGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return mapper.NewSchedules(lessons, calls), nil
}

func (s *ScheduleService) GetScheduleByWeekday(ctx context.Context, groupID uint, weekday, week string) (dto.ScheduleResponse, error) {
	w, ok := weekdays.ParseWeekday(weekday)
	if !ok {
		return dto.ScheduleResponse{}, domain.InvalidDataError{Msg: "invalid weekday"}
	}

	offset, err := s.getWeekOffset(week)
	if err != nil {
		return dto.ScheduleResponse{}, err
	}
	date := weekdays.GetDateByWeekday(w, offset)
	return s.GetScheduleForDate(ctx, groupID, date)
}

func (s *ScheduleService) GetSchedulesByWeek(ctx context.Context, groupID uint, week string) ([]dto.ScheduleResponse, error) {
	offset, err := s.getWeekOffset(week)
	if err != nil {
		return nil, err
	}
	startDate, endDate := weekdays.WeekBounds(offset)
	return s.GetSchedulesForDates(ctx, groupID, startDate, endDate)
}

func (s *ScheduleService) getWeekOffset(week string) (int, error) {
	switch week {
	case "previous":
		return -1, nil
	case "current", "":
		return 0, nil
	case "next":
		return 1, nil
	default:
		return 0, domain.InvalidDataError{Msg: "invalid week (expected previous, current, next)"}
	}
}

func (s *ScheduleService) getCallsByGroup(ctx context.Context, groupID uint) ([]domain.Call, error) {
	college, err := s.collegeRepo.GetByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return s.callRepo.GetByCollegeID(ctx, college.ID)
}
