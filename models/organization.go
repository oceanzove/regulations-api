package models

type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Organization struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetDepartmentOutput struct {
	Departments []Department `json:"departments"`
}

type EmployeeDepartment struct {
	EmployeeID   string `db:"employee_id" json:"employeeId"`
	DepartmentID string `db:"department_id" json:"departmentId"`
}

type GetEmployeeDepartmentOutput struct {
	EmployeeDepartment []EmployeeDepartment `json:"employeeDepartment"`
}

type EmployeePosition struct {
	EmployeeID string `db:"employee_id" json:"employeeId"`
	PositionID string `db:"position_id" json:"positionId"`
}

type GetEmployeePositionOutput struct {
	EmployeePosition []EmployeePosition `json:"employeePosition"`
}

type DepartmentPosition struct {
	DepartmentID string `db:"department_id" json:"departmentId"`
	PositionID   string `db:"position_id" json:"positionId"`
}

type GetDepartmentPositionOutput struct {
	DepartmentPosition []DepartmentPosition `json:"departmentPosition"`
}

type UpdateEmployeeDepartment struct {
	EmployeeID   string `json:"employeeId"`
	DepartmentID string `json:"departmentId"`
}

type Position struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateEmployeePosition struct {
	EmployeeID string `json:"employeeId"`
	PositionID string `json:"positionId"`
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

type CreatePositionInput struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	DepartmentID string `json:"departmentId" db:"department_id"`
}

type UpdatePositionInput struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	DepartmentID string `json:"departmentId" db:"department_id"`
}

type CreateDepartmentInput struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type UpdateDepartmentInput struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
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
