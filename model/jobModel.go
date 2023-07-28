package model

import "time"

type Job struct {
	ID        string    `json:"id"  gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Status    bool      `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

type JobCreation struct {
	ID        string    `json:"id" gorm:"id"`
	Name      string    `json:"name" gorm:"name"`
	Status    bool      `json:"status" gorm:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type JobUpdate struct {
	Name      *string   `json:"name" gorm:"name"`
	Status    *bool     `json:"status" gorm:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Job) TableName() string         { return "job" }
func (JobCreation) TableName() string { return Job{}.TableName() }
func (JobUpdate) TableName() string   { return Job{}.TableName() }
