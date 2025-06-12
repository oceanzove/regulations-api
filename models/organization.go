package models

type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetDepartmentOutput struct {
	Departments []Department `json:"departments"`
}

type Position struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetPositionOutput struct {
	Positions []Position `json:"positions"`
}

type Employee struct {
	ID                 string `json:"id"`
	FullName           string `json:"fullName" db:"full_name"`
	PhoneNumber        string `json:"phoneNumber" db:"phone_number"`
	BirthDate          string `json:"birthDate" db:"birth_date"`
	EmploymentDate     string `json:"employmentDate" db:"employment_date"`
	ResidentialAddress string `json:"residentialAddress" db:"residential_address"`
	MaritalStatus      string `json:"maritalStatus" db:"marital_status"`
	Email              string `json:"email" db:"email"`
}

type CreateEmployeeInput struct {
	Employee     Employee `json:"employee"`
	Account      Account  `json:"account"`
	DepartmentID string   `json:"departmentId"`
	PositionID   string   `json:"positionId"`
}

type GetEmployeesOutput struct {
	Employees []Employee `json:"employees"`
}
