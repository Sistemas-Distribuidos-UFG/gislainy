package main

// Employee is the representation of employee's structure
type Employee struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Salary float64 `json:"salary"`
}
// CalculateReajustSalary reajust the salary according to the employee's position
func CalculateReajustSalary(employee Employee) Employee {
	switch employee.Role {
	case "Operator":
		employee.Salary *= 1.2
	case "Developer":
		employee.Salary *= 1.18
	}
	return employee
}

