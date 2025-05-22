package models

type Step struct {
	ID          string `json:"id"`
	Order       int    `json:"order"`
	Title       string `json:"title"`
	ProcessID   string `json:"processId"`
	Description string `json:"description"`
	Responsible string `json:"responsible"`
}

type CreateStepsInput struct {
	Steps []Step `json:"steps"`
}
