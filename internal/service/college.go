package service

import (
	"context"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/dto"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"
)

type CollegeService struct {
	collegeRepo repository.CollegeRepo
	campusRepo  repository.CampusRepo
}

func NewCollegeService(collegeRepo repository.CollegeRepo, campusRepo repository.CampusRepo) *CollegeService {
	return &CollegeService{collegeRepo: collegeRepo, campusRepo: campusRepo}
}

func (s *CollegeService) GetCollegeByGroupID(ctx context.Context, groupID uint) (dto.CollegeResponse, error) {
	college, err := s.collegeRepo.GetByGroupID(ctx, groupID)
	if err != nil {
		return dto.CollegeResponse{}, err
	}

	campuses, err := s.campusRepo.GetByCollegeID(ctx, college.ID)
	if err != nil {
		return dto.CollegeResponse{}, err
	}

	return college.ToDTO(campuses), nil
}

func (s *CollegeService) GetColleges(ctx context.Context, name string) ([]dto.CollegeResponse, error) {
	var colleges []domain.College
	var err error
	if name != "" {
		colleges, err = s.collegeRepo.GetByName(ctx, name)
	} else {
		colleges, err = s.collegeRepo.GetAll(ctx)
	}
	if err != nil {
		return nil, err
	}

	collegeIDs := make([]uint, len(colleges))
	for i, c := range colleges {
		collegeIDs[i] = c.ID
	}

	campuses, err := s.campusRepo.GetByCollegeIDs(ctx, collegeIDs)
	if err != nil {
		return nil, err
	}

	campusesByCollege := toSliceCampusMap(campuses)

	responses := make([]dto.CollegeResponse, 0, len(colleges))
	for _, college := range colleges {
		responses = append(responses, college.ToDTO(campusesByCollege[college.ID]))
	}

	return responses, nil
}

func (s *CollegeService) GetCollege(ctx context.Context, collegeID uint) (dto.CollegeResponse, error) {
	college, err := s.collegeRepo.Get(ctx, collegeID)
	if err != nil {
		return dto.CollegeResponse{}, err
	}

	campuses, err := s.campusRepo.GetByCollegeID(ctx, collegeID)
	if err != nil {
		return dto.CollegeResponse{}, err
	}

	return college.ToDTO(campuses), nil
}

func toSliceCampusMap(campuses []domain.Campus) map[uint][]domain.Campus {
	result := make(map[uint][]domain.Campus)
	for _, campus := range campuses {
		result[campus.CollegeID] = append(result[campus.CollegeID], campus)
	}
	return result
}
