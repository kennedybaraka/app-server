package validation

import (
	"github.com/gookit/validate"
	"github.com/kennedybaraka/app-server/models"
)

func ValidateCreateOwnerStruct(user models.Owner) error {
	v := validate.Struct(user)

	if !v.Validate() {
		return v.Errors
	}
	return nil
}
