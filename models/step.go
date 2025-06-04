package models

type Step struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Order       int    `json:"order" db:"order"`
	ProcessID   string `json:"processId" db:"process_id"`
	Description string `json:"description" db:"description"`
	Responsible string `json:"responsible" db:"responsible"`
}
type CreateStepsInput struct {
	Steps []Step `json:"steps"`
}
