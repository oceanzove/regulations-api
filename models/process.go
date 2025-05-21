package models

type Process struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetProcessesOutput struct {
	Processes []Process `json:"processes"`
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
