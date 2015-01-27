package gactory

import (
	"reflect"
	"strconv"
	"errors"
	"fmt"
)


// generateValue does the heavy lifting, actually taking the field tag and setting the field
// according to the "gactory" field tag.
func generateValue(value reflect.Value, structField reflect.StructField, structName string) (val reflect.Value, sliceSize int, err error) {
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
		result, err = Fill(value.Interface())

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
