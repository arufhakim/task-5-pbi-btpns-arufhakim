package helpers

import (
	"github.com/asaskevich/govalidator"
)

func ValidateUserInputForAuthentication(userValidation interface{}) (bool, error) {
	result, err := govalidator.ValidateStruct(userValidation)
	if err != nil {
		return result, err
	}

	return result, nil
}
