package gactory

import (
	"reflect"
	"strings"
)

type LinkStruct struct {}

// Make will recursively look through a struct, determine if it contains field tags containing "gactory" and then
// generate a object filled out with fields according the to the "gactory" tag.
func Fill(objectStruct interface{}) (createdObj interface{}, err error) {
	objType := reflect.TypeOf(objectStruct)
	obj := reflect.New(objType).Elem()

	if !containsGactoryTag(obj, objType) {
		return obj.Interface(), nil
	}

	for i := 0; i < obj.NumField(); i++ {
		value := obj.Field(i)
		structField := objType.Field(i)
		if value.CanSet() && value.IsValid() {
			valueToSet, _, err := generateValue(value, structField, objType.Name())
			if err != nil {
				return nil, err
			}
			value.Set( valueToSet )
		} else {
			err := generateError(value.Type().Name(), structField.Name, objType.Name(), "Cannot generate non-settable value")

			if err != nil {
				return nil, err
			}
		}
	}

	return obj.Interface(), nil
}

// hasLinkedStruct loops through the obj and looks for a field with a tag named "GactoryLink".
// If it finds this, it will return true and the type of that field.
func hasLinkedStruct(obj reflect.Value, objType reflect.Type) (isLinked bool, linkType reflect.Type) {
	for i := 0; i < obj.NumField(); i++ {
		structField := objType.Field(i)
		tag := structField.Tag.Get("gactory")

		if ( strings.Contains(tag, "GactoryLink") ) {
			return true, obj.Field(i).Type()
		}
	}
	return false, linkType
}

//containsGactoryTag will determine if a given struct and type contains any "gactory" fields
func containsGactoryTag(obj reflect.Value, objType reflect.Type) (generateTag bool) {
	for i := 0; i < obj.NumField(); i++ {
		structField := objType.Field(i)
		tag := structField.Tag.Get("gactory")
		if tag != "" {
			return true
			break
		}
	}
	return generateTag
}
