package utilities

import (
	ports "github.com/Prompiriya084/go-authen/Internal/Core/Ports/Utilities"
	"github.com/go-playground/validator/v10"
)

type validatorImpl struct {
	validate *validator.Validate
}

func NewValidator() ports.Validator {
	return &validatorImpl{
		validate: validator.New(),
	}
}
func (v *validatorImpl) ValidateStruct(s interface{}) error {
	if err := v.validate.Struct(s); err != nil {
		return err
	}
	return nil
}
