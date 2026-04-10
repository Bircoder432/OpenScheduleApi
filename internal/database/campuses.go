package database

import (
	"context"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/database/models"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"
	"gorm.io/gorm"
)

type CampusDb struct {
	db *gorm.DB
}

func NewCampusDb(db *gorm.DB) *CampusDb {
	return &CampusDb{db: db}
}

func (c CampusDb) CreateMany(ctx context.Context, campuses []domain.Campus) error {
	var campusModels []models.Campus
	for _, campus := range campuses {
		campusModels = append(campusModels, models.Campus{
			Name:      campus.Name,
			CollegeID: campus.CollegeID,
		})
	}
	return c.db.WithContext(ctx).Create(&campusModels).Error
}

func (c CampusDb) GetByID(ctx context.Context, campusID uint) (domain.Campus, error) {
	var campusModel models.Campus
	err := c.db.WithContext(ctx).First(&campusModel, campusID).Error
	if err == gorm.ErrRecordNotFound {
		return domain.Campus{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.Campus{}, err
	}
	return campusModel.ToDomain(), nil
}

func (c CampusDb) GetByName(ctx context.Context, collegeId uint, name string) ([]domain.Campus, error) {
	var campusModels []models.Campus
	err := c.db.WithContext(ctx).Where(
		models.Campus{CollegeID: collegeId, Name: name},
	).Find(&campusModels).Error
	if err != nil {
		return nil, err
	}
	var campuses []domain.Campus
	for _, campusModel := range campusModels {
		campuses = append(campuses, campusModel.ToDomain())
	}
	return campuses, nil
}

func (c CampusDb) GetByCollegeID(ctx context.Context, collegeID uint) ([]domain.Campus, error) {
	var campusModels []models.Campus
	if err := c.db.WithContext(ctx).Where("college_id = ?", collegeID).
		Find(&campusModels).Error; err != nil {
		return nil, err
	}
	var campuses []domain.Campus
	for _, campusModel := range campusModels {
		campuses = append(campuses, campusModel.ToDomain())
	}
	return campuses, nil
}

func (c CampusDb) GetByCollegeIDs(ctx context.Context, collegeIDs []uint) ([]domain.Campus, error) {
	var campusModels []models.Campus
	err := c.db.WithContext(ctx).Where("college_id IN ?", collegeIDs).Find(&campusModels).Error
	if err != nil {
		return nil, err
	}
	var campuses []domain.Campus
	for _, campusModel := range campusModels {
		campuses = append(campuses, campusModel.ToDomain())
	}
	return campuses, nil
}

func (c CampusDb) WithTx(tx repository.Tx) repository.CampusRepo {
	return CampusDb{db: tx.DB().(*gorm.DB)}
}
