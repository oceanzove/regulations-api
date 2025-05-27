package models

import "time"

type Process struct {
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Responsible string    `json:"responsible" db:"responsible"`
	Description string    `json:"description" db:"description"`
	AccountID   string    `json:"account_id" db:"account_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type GetProcessesOutput struct {
	Processes []Process `json:"processes"`
}

type LinkRegulationToProcessInput struct {
	ProcessID    string `json:"process_id" binding:"required,uuid"`
	RegulationID string `json:"regulation_id" binding:"required,uuid"`
}

type UpdateProcessInput struct {
	ID          string
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateProcessInput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateProcessOutput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
