package database

import (
	"context"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/database/models"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/domain"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CallDb struct {
	db *gorm.DB
}

func NewCallDb(db *gorm.DB) *CallDb {
	return &CallDb{db: db}
}

func (c CallDb) GetByCollegeID(ctx context.Context, collegeID uint) ([]domain.Call, error) {
	var calls []models.Call
	var domainCalls []domain.Call

	if err := c.db.WithContext(ctx).Where("college_id = ?", collegeID).Find(&calls).Error; err != nil {
		return nil, err
	}
	for _, call := range calls {
		domainCalls = append(domainCalls, call.ToDomain())
	}
	return domainCalls, nil
}

func (c CallDb) UpsertMany(ctx context.Context, calls []domain.Call) error {
	var modelsCalls []models.Call

	for _, call := range calls {
		modelsCalls = append(modelsCalls, models.Call{
			CallID:  call.ID,
			Weekday: call.Weekday,
			Begins: datatypes.NewTime(
				call.Begins.Hour(), call.Begins.Minute(), 0, 0,
			),
			Ends: datatypes.NewTime(
				call.Ends.Hour(), call.Ends.Minute(), 0, 0,
			),
			Order:     call.Order,
			CollegeID: call.CollegeID,
		})
	}
	return c.db.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "college_id"}, {Name: "weekday"}, {Name: "order"}},
			DoUpdates: clause.AssignmentColumns([]string{"begins", "ends"}),
		},
	).Create(&modelsCalls).Error
}
