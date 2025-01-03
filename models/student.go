package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}
type Student struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstName" gorm:"not null" validate:"required,min=3"`
	LastName  string `json:"lastName" gorm:"not null" validate:"required,min=3"`
}

func (s *Student) BeforeCreate(dbInstance *gorm.DB) (err error){
	if error := validate.Struct(s); error != nil {
		var validationErrors []string
		for _, e := range error.(validator.ValidationErrors){
			if e.Tag() == "min" && e.Field() == "FirstName" {
				validationErrors = append(validationErrors, "firstName should be atleast 3 character long")
			} else if e.Tag() == "min" && e.Field() == "LastName"{
				validationErrors = append(validationErrors, "lastName should be atleast 3 character long")
			} else {
				validationErrors = append(validationErrors, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))	
			}
		}
	return fmt.Errorf("validation failed: %v", validationErrors)
	}
	return nil
}