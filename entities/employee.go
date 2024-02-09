package entities

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name" validate:"required"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"-"`
}
