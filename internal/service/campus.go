package service

import (
	"context"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/dto"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/utils"
)

type CampusService struct {
	collegeRepo repository.CollegeRepo
	campusRepo  repository.CampusRepo
	groupRepo   repository.StudentGroupRepo
}

func NewCampusService(
	campusRepo repository.CampusRepo,
	groupRepo repository.StudentGroupRepo,
	collegeRepo repository.CollegeRepo) *CampusService {
	return &CampusService{campusRepo: campusRepo, groupRepo: groupRepo, collegeRepo: collegeRepo}
}

func (s *CampusService) GetCampusesByCollegeID(ctx context.Context, collegeID uint, name string) ([]dto.CampusResponse, error) {
	if _, err := s.collegeRepo.Get(ctx, collegeID); err != nil {
		return nil, err
	}
	var campuses []domain.Campus
	var err error
	if name != "" {
		campuses, err = s.campusRepo.GetByName(ctx, collegeID, name)
	} else {
		campuses, err = s.campusRepo.GetByCollegeID(ctx, collegeID)
	}
	if err != nil {
		return nil, err
	}

	campusIds := utils.ToNewSlice(campuses,
		func(campus domain.Campus) (uint, bool) { return campus.ID, true })
	studentGroups, err := s.groupRepo.GetByCampusIDs(ctx, campusIds)
	if err != nil {
		return nil, err
	}

	studentGroupsByCampus := toSliceGroupMap(studentGroups)

	responses := make([]dto.CampusResponse, 0, len(campuses))
	for _, campus := range campuses {
		responses = append(responses, campus.ToDTO(studentGroupsByCampus[campus.ID]))
	}
	return responses, nil
}

func (s *CampusService) GetCampusByID(ctx context.Context, ID uint) (dto.CampusResponse, error) {
	campus, err := s.campusRepo.GetByID(ctx, ID)
	if err != nil {
		return dto.CampusResponse{}, err
	}
	studentGroups, err := s.groupRepo.GetByCampusID(ctx, ID)
	if err != nil {
		return dto.CampusResponse{}, err
	}
	return campus.ToDTO(studentGroups), nil
}

func toSliceGroupMap(groups []domain.StudentGroup) map[uint][]domain.StudentGroup {
	result := make(map[uint][]domain.StudentGroup)
	for _, group := range groups {
		result[group.CampusID] = append(result[group.CampusID], group)
	}
	return result
}
