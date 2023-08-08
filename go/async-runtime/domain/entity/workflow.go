package entity

import (
	"time"

	"gorm.io/datatypes"
)

type WorkflowEntity struct {
	ID         uint           `gorm:"primaryKey"`
	ResourceId string         `gorm:"column:rid;index"`
	Type       string         `gorm:"column:type;index"`
	Status     string         `gorm:"column:status;index"`
	Payload    datatypes.JSON `gorm:"column:payload"`
	CreatedAt  time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
}

func (WorkflowEntity) TableName() string {
	return "go.workflow"
}

type WorkflowStateEntity struct {
	ID         uint      `gorm:"primaryKey"`
	WorkflowId uint      `gorm:"workflow_id"`
	Name       string    `gorm:"column:name;index"`
	Status     string    `gorm:"column:status;index"`
	CreatedAt  time.Time `gorm:"index"`
	UpdatedAt  time.Time `gorm:"index"`
}

func (WorkflowStateEntity) TableName() string {
	return "go.workflow_state"
}

type WorkflowVariablesEntity struct {
	ID         uint           `gorm:"primaryKey"`
	WorkflowId uint           `gorm:"workflow_id"`
	Variables  datatypes.JSON `gorm:"column:variables"`
	LastTask   string         `gorm:"column:last_task;index"`
	CreatedAt  time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
}

func (WorkflowVariablesEntity) TableName() string {
	return "go.workflow_variable"
}
