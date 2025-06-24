package models

type Regulation struct {
	ID           string `json:"id" db:"id"`
	Title        string `json:"title" db:"title"`
	Content      string `json:"content" db:"content"`
	DepartmentID string `json:"departmentId" db:"department_id"`
}

type LinkSectionToRegulation struct {
	ID           string `json:"id" binding:"required,uuid"`
	SectionID    string `json:"sectionId" binding:"required,uuid"`
	RegulationID string `json:"regulationId"`
	Order        int    `json:"order"`
}

type Section struct {
	ID        string `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
	AccountID string `json:"account_id" db:"account_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

type GetRegulationsOutput struct {
	Regulations []Regulation `json:"regulations"`
}

type GetSectionsOutput struct {
	Sections []Section `json:"sections"`
}

type GetSectionByRegulationOutput struct {
	SectionIDs []string `json:"sectionsIds"`
}

type UpdateRegulationInput struct {
	ID      string
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateRegulationInput struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateSectionInput struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateRegulationOutput struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
