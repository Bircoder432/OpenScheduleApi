package database

import (
	"context"
	"errors"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/database/models"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"

	"gorm.io/gorm"
)

type CollegeDb struct {
	db *gorm.DB
}

func NewCollegeDb(db *gorm.DB) *CollegeDb {
	return &CollegeDb{db: db}
}

func (c CollegeDb) Create(ctx context.Context, college domain.College) (uint, error) {
	collegeModel := models.College{Name: college.Name, Token: college.Token}
	return collegeModel.CollegeID, c.db.WithContext(ctx).Create(&collegeModel).Error
}

func (c CollegeDb) Get(ctx context.Context, collegeID uint) (domain.College, error) {
	var collegeModel models.College
	err := c.db.WithContext(ctx).First(&collegeModel, collegeID).Error
	if err == gorm.ErrRecordNotFound {
		return domain.College{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.College{}, err
	}
	return collegeModel.ToDomain(), nil
}

func (c CollegeDb) GetByName(ctx context.Context, name string) ([]domain.College, error) {
	var collegeModels []models.College
	err := c.db.WithContext(ctx).Where("name = ?", name).Find(&collegeModels).Error
	if err != nil {
		return nil, err
	}
	var colleges []domain.College
	for _, collegeModel := range collegeModels {
		colleges = append(colleges, collegeModel.ToDomain())
	}
	return colleges, nil
}

func (c CollegeDb) GetAll(ctx context.Context) ([]domain.College, error) {
	var collegeModels []models.College
	err := c.db.WithContext(ctx).Find(&collegeModels).Error
	if err != nil {
		return nil, err
	}
	var colleges []domain.College
	for _, collegeModel := range collegeModels {
		colleges = append(colleges, collegeModel.ToDomain())
	}
	return colleges, nil
}

func (c CollegeDb) GetByGroupID(ctx context.Context, groupID uint) (domain.College, error) {
	var collegeModel models.College
	err := c.db.WithContext(ctx).Table("student_groups AS sg").
		Select("c.college_id, c.name").
		Joins("JOIN campuses AS c ON c.campus_id = sg.campus_id").
		Joins("JOIN colleges AS co ON co.college_id = c.college_id").
		Where("sg.student_group_id = ?", groupID).
		Scan(&collegeModel).Error
	if err == gorm.ErrRecordNotFound {
		return domain.College{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.College{}, err
	}
	return collegeModel.ToDomain(), nil
}

func (c CollegeDb) GetByToken(ctx context.Context, token string) (domain.College, error) {
	var college models.College
	if err := c.db.WithContext(ctx).Where("token = ?", token).First(&college).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.College{}, domain.ErrNotFound
	} else if err != nil {
		return domain.College{}, err
	}
	return college.ToDomain(), nil
}

func (c CollegeDb) Delete(ctx context.Context, collegeID uint) error {
	result := c.db.WithContext(ctx).Delete(&models.College{}, collegeID)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (c CollegeDb) WithTx(tx repository.Tx) repository.CollegeRepo {
	return CollegeDb{db: tx.DB().(*gorm.DB)}
}
