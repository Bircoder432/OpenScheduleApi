package database

import (
	"context"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/database/models"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GroupDb struct {
	db *gorm.DB
}

func NewGroupDb(db *gorm.DB) *GroupDb {
	return &GroupDb{db: db}
}

func (g GroupDb) UpsertMany(ctx context.Context, groups []domain.StudentGroup) error {
	var groupModels = []models.StudentGroup{}
	for _, group := range groups {
		groupModels = append(groupModels, models.StudentGroup{
			Name:     group.Name,
			CampusID: group.CampusID,
		})
	}
	return g.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).Create(&groupModels).Error
}

func (g GroupDb) GetByID(ctx context.Context, groupID uint) (domain.StudentGroup, error) {
	var group models.StudentGroup
	if err := g.db.WithContext(ctx).First(&group, groupID).Error; err != nil {
		return domain.StudentGroup{}, err
	}
	return group.ToDomain(), nil
}

func (g GroupDb) GetByCampusID(ctx context.Context, campusID uint) ([]domain.StudentGroup, error) {
	var groupModels []models.StudentGroup
	err := g.db.WithContext(ctx).Where("campus_id = ?", campusID).Find(&groupModels).Error
	if err != nil {
		return nil, err
	}
	var groups []domain.StudentGroup
	for _, groupModel := range groupModels {
		groups = append(groups, groupModel.ToDomain())
	}
	return groups, nil
}

func (g GroupDb) GetByCampusIDs(ctx context.Context, campusIDs []uint) ([]domain.StudentGroup, error) {
	var groupModels []models.StudentGroup
	err := g.db.WithContext(ctx).Where("campus_id IN ?", campusIDs).Find(&groupModels).Error
	if err != nil {
		return nil, err
	}
	var groups []domain.StudentGroup
	for _, groupModel := range groupModels {
		groups = append(groups, groupModel.ToDomain())
	}
	return groups, nil
}
func (g GroupDb) GetByCampusIDAndName(ctx context.Context, campusID uint, name string) ([]domain.StudentGroup, error) {
	var groupModels []models.StudentGroup
	err := g.db.WithContext(ctx).Where("campus_id = ?", campusID).Where("name = ?", name).Find(&groupModels).Error
	if err != nil {
		return nil, err
	}
	var groups []domain.StudentGroup
	for _, groupModel := range groupModels {
		groups = append(groups, groupModel.ToDomain())
	}
	return groups, nil
}
func (g GroupDb) GetByCollegeIDAndName(ctx context.Context, collegeID uint, name string) ([]domain.StudentGroup, error) {
	var groupModels []models.StudentGroup
	err := g.db.WithContext(ctx).Where("college_id = ?", collegeID).Where("name = ?", name).Find(&groupModels).Error
	if err != nil {
		return nil, err
	}
	var groups []domain.StudentGroup
	for _, groupModel := range groupModels {
		groups = append(groups, groupModel.ToDomain())
	}
	return groups, nil
}
func (g GroupDb) GetByCollegeID(ctx context.Context, collegeID uint) ([]domain.StudentGroup, error) {
	var groupModels []models.StudentGroup
	if err := g.db.WithContext(ctx).
		Joins("JOIN campuses ON campuses.campus_id = student_groups.campus_id").
		Where("campuses.college_id = ?", collegeID).
		Find(&groupModels).Error; err != nil {
		return nil, err
	}
	var groups []domain.StudentGroup
	for _, groupModel := range groupModels {
		groups = append(groups, groupModel.ToDomain())
	}
	return groups, nil
}
