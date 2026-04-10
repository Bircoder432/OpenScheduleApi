package service

import (
	"context"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/dto"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"
)

type StudentGroupService struct {
	groupRepo   repository.StudentGroupRepo
	campusRepo  repository.CampusRepo
	collegeRepo repository.CollegeRepo
}

func NewStudentGroupService(
	groupRepo repository.StudentGroupRepo,
	campusRepo repository.CampusRepo,
	collegeRepo repository.CollegeRepo) *StudentGroupService {
	return &StudentGroupService{groupRepo: groupRepo, campusRepo: campusRepo, collegeRepo: collegeRepo}
}

func (s *StudentGroupService) GetGroups(ctx context.Context, campusID uint, name string) ([]dto.StudentGroupResponse, error) {
	if _, err := s.campusRepo.GetByID(ctx, campusID); err != nil {
		return nil, err
	}
	var groups []domain.StudentGroup
	var err error
	if name != "" {
		groups, err = s.groupRepo.GetByCampusIDAndName(ctx, campusID, name)
	} else {
		groups, err = s.groupRepo.GetByCampusID(ctx, campusID)
	}
	if err != nil {
		return nil, err
	}

	responses := make([]dto.StudentGroupResponse, 0, len(groups))
	for _, group := range groups {
		responses = append(responses, group.ToDTO())
	}
	return responses, nil
}

func (s *StudentGroupService) GetGroupsByCollegeID(ctx context.Context, collegeID uint, name string) ([]dto.StudentGroupResponse, error) {
	if _, err := s.collegeRepo.Get(ctx, collegeID); err != nil {
		return nil, err
	}
	var groups []domain.StudentGroup
	var err error
	if name != "" {
		groups, err = s.groupRepo.GetByCollegeIDAndName(ctx, collegeID, name)
	} else {
		groups, err = s.groupRepo.GetByCollegeID(ctx, collegeID)
	}
	if err != nil {
		return nil, err
	}

	responses := make([]dto.StudentGroupResponse, 0, len(groups))
	for _, group := range groups {
		responses = append(responses, group.ToDTO())
	}
	return responses, nil
}

func (s *StudentGroupService) GetGroup(ctx context.Context, ID uint) (dto.StudentGroupResponse, error) {
	group, err := s.groupRepo.GetByID(ctx, ID)
	if err != nil {
		return dto.StudentGroupResponse{}, err
	}
	return group.ToDTO(), nil
}
