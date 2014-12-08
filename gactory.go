package gactory

import (
	"log"
	"reflect"
	"strconv"
	"errors"
	"fmt"
)

func GenMock(objectStruct interface{}) (interface{}, []error) {
	objType := reflect.TypeOf(objectStruct)
	obj := reflect.New(objType).Elem()
	errs := []error{}
	for i := 0; i < obj.NumField(); i++ {
		value := obj.Field(i)
		structField := objType.Field(i)
		if ( value.CanSet() && value.IsValid() ) {
			valueToSet, makeErrs := MakeValue(value, structField, objType.Name())
			errs = append(errs, makeErrs...)
			value.Set(reflect.ValueOf(valueToSet) )
		} else {
			err := generateError(value.Type().Name(), structField.Name, objType.Name(), "Cannot generate non-settable value")
			errs = append(errs, err)
		}
	}
	return obj.Interface(), errs
}

func MakeValue(value reflect.Value, structField reflect.StructField, structName string) (interface{}, []error)  {
	genKind := value.Type().Kind()
	tag := structField.Tag.Get("gactory")
	log.Printf("tag %v genKind %v",tag, genKind)
	if (tag == "") {
		err := generateError(value.Type().Name(), structField.Name, structName, "Empty tag found")
		return nil, []error{ err }
	}
	result := *new(interface{})
	var err error
	switch (genKind){
	case reflect.Bool:
		result, err = strconv.ParseBool(tag)
	case reflect.Float32:
		result, err = strconv.ParseFloat(tag, 32)
	case reflect.Float64:
		result, err = strconv.ParseFloat(tag, 64)
	case reflect.Int:
		intObj, err := strconv.ParseInt(tag, 10, 64)
		result, err = int(intObj), err
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
		return GenMock(value.Interface())
	case reflect.Array:
		//@TODO: implement this
	default:
		err = generateError(value.Type().Name(), structField.Name, structName, "Could not find type")
		return nil, []error{err}
	}
	if (err == nil) {
		return result, []error{}
	} else {
		return nil, []error{err}
	}
}

func generateError(genType, field, structName, msg string) error {
	errorMsg := fmt.Sprintf("%v occurred when generating an obj of type '%v' in field '%v' on struct '%v'", msg, genType, field, structName)
	return errors.New(errorMsg)
}

