package models

type Employees struct {
	ID                 string              `json:"id" validate:"omitempty,uuid"`
	Name               string              `json:"name" validate:"required,alpha"`
	Phone              string              `json:"phone" validate:"e164"`
	Email              string              `json:"email" validate:"required,email"`
	Technologies       []string            `json:"technologies" validate:"required,max=10,min=1"`
	Savings            int64               `json:"savings"`
	CompanyInformation *CompanyInformation `json:"companyinformation"`
	Address            *Address            `json:"Address"`
}

type CompanyInformation struct {
	Company_Name     string `json:"companyname" validate:"required,alpha"`
	Company_Location string `json:"company_location"  validate:"required,alpha"`
	Position         string `json:"position"  validate:"required"`
	Salary           int64  `json:"salary"  validate:"required"`
}

type Address struct {
	City    string `json:"city"  validate:"required,alpha"`
	State   string `json:"state"  validate:"required,alpha"`
	Country string `json:"country"  validate:"required,alpha"`
}

type Expenditure struct {
	Rent  int64 `json:"rent"`
	Bills int64 `json:"bills"`
	Food  int64 `json:"food"`
}

type Conf struct {
	Summary string `json:"summary"`
	Port    string `json:"port"`
}
