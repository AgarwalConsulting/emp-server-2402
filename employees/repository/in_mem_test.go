package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	existingEmp, err := sut.ListAll()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(existingEmp))

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			newEmp := entities.Employee{Name: "Gaurav"}

			_, err := sut.Save(newEmp)
			assert.Nil(t, err)

			_, err = sut.ListAll()
			assert.Nil(t, err)
		}()
	}

	wg.Wait()

	employees, err := sut.ListAll()
	assert.Equal(t, 102, len(employees))
}
