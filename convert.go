package gactory

import "reflect"

// Convert will take any struct value passed into it and check for a field tag with the name "GactoryLink"
// if it finds this field tag, it'll copy the obj passed in to new struct with the field's type
func Convert(objStruct interface{}) (createdObj interface{}, err error) {
	objType := reflect.TypeOf(objStruct)
	objVal := reflect.ValueOf(objStruct)

	isLinked, linkType := hasLinkedStruct(objVal, objType)

	linkObj := reflect.New(linkType).Elem()
	if isLinked {
		for i := 0; i < objType.NumField(); i++ {
			copyFieldName := objType.Field(i).Name
			fieldToCopy := objVal.FieldByName(copyFieldName)
			fieldToCopyTo := linkObj.FieldByName(copyFieldName)
			if (fieldToCopyTo.IsValid()) {
				fieldToCopyTo.Set(fieldToCopy)
			}
		}
	}

	return linkObj.Interface(), nil
}
