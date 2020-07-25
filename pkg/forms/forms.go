package forms

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/vsokoltsov/users-service/pkg/models"
)

type baseForm interface {
	Submit() (models.DBModel, *map[string][]string)
	validate() map[string][]string
}

func DefaultFormValidation(bf baseForm) map[string][]string {
	var (
		valid   = validator.New()
		errsMap = make(map[string][]string)
	)
	fieldValues, fieldTags := GetFieldsWithValues(bf)
	for key, value := range fieldValues {
		tag := fieldTags[key]
		ferr := valid.Var(value, tag.(string))
		if ferr != nil {
			var errStrings []string
			errsData := ferr.(validator.ValidationErrors)
			for _, errItem := range errsData {
				errStrings = append(errStrings, errItem.Tag())
			}
			errsMap[key] = errStrings
		}
	}
	return errsMap
}

// GetFieldsWithValues return maps of fields and tags
func GetFieldsWithValues(bf baseForm) (map[string]interface{}, map[string]interface{}) {
	var (
		fieldValue = make(map[string]interface{})
		fieldTag   = make(map[string]interface{})
	)
	rfields := reflect.TypeOf(bf).Elem()
	rvalues := reflect.ValueOf(bf).Elem()
	for i := 0; i < rfields.NumField(); i++ {
		field := rfields.Field(i)
		value := rvalues.Field(i)
		tag := field.Tag.Get("validate")
		fieldValue[field.Name] = value.Interface()
		fieldTag[field.Name] = tag
	}
	return fieldValue, fieldTag
}
