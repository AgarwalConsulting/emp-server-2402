package repository

import "algogrit.com/emp-server/entities"

type inmem struct {
	empList []entities.Employee
}

func (repo *inmem) ListAll() ([]entities.Employee, error) {
	return repo.empList, nil
}

func (repo *inmem) Save(newEmp entities.Employee) (*entities.Employee, error) {
	newEmp.ID = len(repo.empList) + 1

	repo.empList = append(repo.empList, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Balaji", "Cloud", 10002},
	}

	return &inmem{empList: employees}
}
