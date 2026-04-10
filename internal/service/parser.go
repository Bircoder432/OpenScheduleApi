package service

import (
	"context"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"
)

type ParserService struct {
	callRepo    repository.CallRepo
	groupRepo   repository.StudentGroupRepo
	lessonRepo  repository.LessonRepo
	campusRepo  repository.CampusRepo
	collegeRepo repository.CollegeRepo
}

func NewParserService(
	callRepo repository.CallRepo,
	groupRepo repository.StudentGroupRepo,
	lessonRepo repository.LessonRepo,
	campusRepo repository.CampusRepo,
	collegeRepo repository.CollegeRepo,
) *ParserService {
	return &ParserService{
		callRepo:    callRepo,
		groupRepo:   groupRepo,
		lessonRepo:  lessonRepo,
		campusRepo:  campusRepo,
		collegeRepo: collegeRepo,
	}
}

func (s *ParserService) GetByToken(ctx context.Context, token string) (domain.College, error) {
	return s.collegeRepo.GetByToken(ctx, token)
}

func (s *ParserService) UpdateGroups(ctx context.Context, campusID uint, groupNames []string) error {
	if _, err := s.campusRepo.GetByID(ctx, campusID); err != nil {
		return err
	}

	groups := make([]domain.StudentGroup, 0, len(groupNames))
	for _, groupName := range groupNames {
		groups = append(groups, domain.StudentGroup{Name: groupName, CampusID: campusID})
	}
	return s.groupRepo.UpsertMany(ctx, groups)
}

func (s *ParserService) UpdateCalls(ctx context.Context, collegeID uint, calls []domain.Call) error {
	if _, err := s.collegeRepo.Get(ctx, collegeID); err != nil {
		return err
	}
	for i := range calls {
		calls[i].CollegeID = collegeID
	}
	return s.callRepo.UpsertMany(ctx, calls)
}

func (s *ParserService) AddLessons(ctx context.Context, lessons []domain.Lesson) error {
	return s.lessonRepo.Add(ctx, lessons)
}
