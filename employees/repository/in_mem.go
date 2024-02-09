package repository

import (
	"sync"

	"algogrit.com/emp-server/entities"
)

type inmem struct {
	empList []entities.Employee
	mut     sync.RWMutex // Default Value: ?
}

func (repo *inmem) ListAll() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	return repo.empList, nil
}

func (repo *inmem) Save(newEmp entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()

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
