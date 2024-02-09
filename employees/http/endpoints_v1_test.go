package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	expectedEmployee := entities.Employee{Name: "Gaurav", Department: "LnD"}
	mockSvc.EXPECT().Create(expectedEmployee).Return(&expectedEmployee, nil)

	jsonBody := `{"name": "Gaurav", "speciality": "LnD"}` // string
	reqBody := strings.NewReader(jsonBody)

	// req := httptest.NewRequest("GET", "/v1/employees", nil)
	req := httptest.NewRequest("POST", "/v1/employees", reqBody)
	respRec := httptest.NewRecorder()

	// sut.CreateV1(respRec, req)
	sut.ServeHTTP(respRec, req)

	resp := respRec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var actualEmp entities.Employee
	json.NewDecoder(resp.Body).Decode(&actualEmp)

	assert.Equal(t, expectedEmployee, actualEmp)
}
