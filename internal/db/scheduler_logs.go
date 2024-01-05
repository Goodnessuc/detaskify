package db

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type ScheduleExecutionLog struct {
	gorm.Model
	ScheduleID uuid.UUID `gorm:"index;not null"`
	ExecutedAt time.Time `gorm:"type:datetime;not null"`
	Status     string    `gorm:"type:varchar(50)"`
	Output     string    `gorm:"type:text"`
}

func (d *Database) LogExecution(ctx context.Context, taskLog *ScheduleExecutionLog) error {
	err := d.Client.WithContext(ctx).Create(taskLog).Error
	if err != nil {
		log.Printf("Error logging execution: %s", err.Error())
		return err
	}
	return nil
}

func (d *Database) GetExecutionLog(ctx context.Context, logID uint) (*ScheduleExecutionLog, error) {
	var taskLog ScheduleExecutionLog
	err := d.Client.WithContext(ctx).Where("id = ?", logID).First(&taskLog).Error
	if err != nil {
		log.Printf("Error retrieving execution taskLog: %s", err.Error())
		return nil, err
	}
	return &taskLog, nil
}

func (d *Database) UpdateExecutionLog(ctx context.Context, logID uint, updatedLog *ScheduleExecutionLog) error {
	err := d.Client.WithContext(ctx).Model(&ScheduleExecutionLog{}).Where("id = ?", logID).Updates(updatedLog).Error
	if err != nil {
		log.Printf("Error updating execution log: %s", err.Error())
		return err
	}
	return nil
}

func (d *Database) DeleteExecutionLog(ctx context.Context, logID uint) error {
	result := d.Client.WithContext(ctx).Where("id = ?", logID).Delete(&ScheduleExecutionLog{})
	if result.Error != nil {
		log.Printf("Error deleting execution log: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (d *Database) ListExecutionLogsByScheduleID(ctx context.Context, scheduleID uuid.UUID) ([]ScheduleExecutionLog, error) {
	var logs []ScheduleExecutionLog
	err := d.Client.WithContext(ctx).Where("schedule_id = ?", scheduleID).Find(&logs).Error
	if err != nil {
		log.Printf("Error listing execution logs for schedule ID %s: %s", scheduleID, err.Error())
		return nil, err
	}
	return logs, nil
}
