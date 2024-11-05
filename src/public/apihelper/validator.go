package apihelper

import (
	"errors"
	"fmt"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
)

// TagMessage tag message for struct validate
const TagMessage = "message"

// TagErrorCode tag code for struct validation in return if the field got error
const TagErrorCode = "errorCode"

type CustomValidate struct {
	*validator.Validate
	Message *string
}

func (customValidate *CustomValidate) init(validate *validator.Validate) {
	customValidate.Validate = validate
}

// Struct validate struct
func (customValidate *CustomValidate) Struct(current interface{}) error {
	errValidate := customValidate.Validate.Struct(current)
	if errValidate != nil {
		for _, err := range errValidate.(validator.ValidationErrors) {
			t := reflect.TypeOf(current)
			for i := 0; i < t.NumField(); i++ {
				if string(t.Field(i).Name) == err.Field() {
					errMsg := t.Field(i).Tag.Get(TagMessage)
					errCode := t.Field(i).Tag.Get(TagErrorCode)
					if len(errMsg) > 0 {
						return errors.New(errCode)
					}
				}
			}
			return fmt.Errorf(strconv.Itoa(common.GeneralBadRequest))
		}
	}
	return nil
}

// TSCustomValidator custom validator
func TSCustomValidator() *CustomValidate {
	customValidate := &CustomValidate{}
	customValidate.init(validator.New())
	return customValidate
}
