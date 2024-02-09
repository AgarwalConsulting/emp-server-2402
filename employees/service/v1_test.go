package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository.NewMockEmployeeRepository(ctrl)
	sut := service.NewV1(mockRepo)

	expectedEmployees := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	// Setup Expections
	mockRepo.EXPECT().ListAll().Return(expectedEmployees, nil)

	employees, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, employees)
	assert.Equal(t, expectedEmployees, employees)
}
