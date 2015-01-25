package gactory

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type LinkStruct struct {}

// Make will recursively look through a struct, determine if it contains field tags containing "gactory" and then
// generate a object filled out with fields according the to the "gactory" tag.
func Make(objectStruct interface{}) (createdObj interface{}, err error) {
	objType := reflect.TypeOf(objectStruct)
	obj := reflect.New(objType).Elem()

	if !containsGactoryTag(obj, objType) {
		return obj.Interface(), nil
	}

	for i := 0; i < obj.NumField(); i++ {
		value := obj.Field(i)
		structField := objType.Field(i)
		if value.CanSet() && value.IsValid() {
			valueToSet, _, err := GenerateValue(value, structField, objType.Name())
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

// GenerateValue does the heavy lifting, actually taking the field tag and setting the field
// according to the "gactory" field tag.
func GenerateValue(value reflect.Value, structField reflect.StructField, structName string) (val reflect.Value, sliceSize int, err error) {
	var result interface{}
	kind := value.Type().Kind()
	tag := structField.Tag.Get("gactory")

	if tag == "" && kind != reflect.Struct {
		return val, sliceSize, generateError(value.Type().Name(), structField.Name, structName, "Empty tag found")
	}

	switch kind {
	case reflect.Bool:
		result, err = strconv.ParseBool(tag)

	case reflect.Float32:
		result, err = strconv.ParseFloat(tag, 32)

	case reflect.Float64:
		result, err = strconv.ParseFloat(tag, 64)

	case reflect.Int:
		intObj := int64(0)
		intObj, err = strconv.ParseInt(tag, 10, 0)
		result = int(intObj)

	case reflect.Int8:
		result, err = strconv.ParseInt(tag, 10, 8)

	case reflect.Int16:
		result, err = strconv.ParseInt(tag, 10, 16)

	case reflect.Int32:
		result, err = strconv.ParseInt(tag, 10, 32)

	case reflect.Int64:
		result, err = strconv.ParseInt(tag, 10, 64)

	case reflect.String:
		result, err = tag, nil

	case reflect.Struct:
		result, err = Make(value.Interface())

	case reflect.Slice:
		result, sliceSize, err = generateSlice(tag, value)

	case reflect.Map:
		err = errors.New("Cannot generate maps")
	default:
		err = generateError(value.Type().Name(), structField.Name, structName, "Could not find type")
		return val, sliceSize, err
	}
	return reflect.ValueOf(result), sliceSize, err
}

func generateError(genType, field, structName, msg string) error {
	errorMsg := fmt.Sprintf("%v occurred when generating an obj of type '%v' in field '%v' on struct '%v'", msg, genType, field, structName)
	return errors.New(errorMsg)
}
