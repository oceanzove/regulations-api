package models

type Regulation struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetRegulationsOutput struct {
	Regulations []Regulation `json:"regulations"`
}

type UpdateRegulationInput struct {
	ID      string
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateRegulationOutput struct {
	ID      string
	Title   string `json:"title"`
	Content string `json:"content"`
}
