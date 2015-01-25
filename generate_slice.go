package gactory

import (
	"strconv"
	"strings"
	"reflect"
	"errors"
)

var typeOfByte = reflect.TypeOf([]byte(nil))
var typeOfString = reflect.TypeOf([]string(nil))
var typeOfInt = reflect.TypeOf([]int(nil))
var typeOfInt8 = reflect.TypeOf([]int8(nil))
var typeOfInt16 = reflect.TypeOf([]int16(nil))
var typeOfInt32 = reflect.TypeOf([]int32(nil))
var typeOfInt64 = reflect.TypeOf([]int64(nil))
var typeOfFloat64 = reflect.TypeOf([]float64(nil))
var typeOfFloat32 = reflect.TypeOf([]float32(nil))
var typeOfBool = reflect.TypeOf([]bool(nil))
var typeOfUint = reflect.TypeOf([]uint(nil))

func generateSlice(tag string, value reflect.Value) (slice interface{}, numVals int, err error) {
	rawValues := strings.Split(tag, ",")

	switch value.Type() {
	case typeOfBool:
		slice, err = generateBoolSlice(rawValues)
	case typeOfUint:
		slice, err = generateUintSlice(rawValues, 0)
	case typeOfByte:
		slice, err = []byte(tag), nil
	case typeOfString:
		slice, err = rawValues, nil
	case typeOfInt:
		slice, err = generateIntSlice(rawValues, 0)
	case typeOfInt8:
		slice, err = generateIntSlice(rawValues, 8)
	case typeOfInt16:
		slice, err = generateIntSlice(rawValues, 16)
	case typeOfInt32:
		slice, err = generateIntSlice(rawValues, 32)
	case typeOfInt64:
		slice, err = generateIntSlice(rawValues, 64)
	case typeOfFloat64:
		slice, err = generateFloatSlice(rawValues, 64)
	case typeOfFloat32:
		slice, err = generateFloatSlice(rawValues, 32)
	default:
		return
	}
	return slice, len(rawValues), err
}

func generateStrSlice(rawValues []string) (slice interface{}, err error){
	strSlice := make([]string, len(rawValues))
	for index, rawValue := range rawValues {
		strSlice[index] = rawValue
	}
	return strSlice, err
}

func generateBoolSlice(rawValues []string) (slice interface{}, err error){
	boolSlice := make([]bool, len(rawValues))
	for index, rawValue := range rawValues {
		boolSlice[index], err = strconv.ParseBool(rawValue)
		if (err != nil) {
			return boolSlice, err
		}
	}
	return boolSlice, err
}

func generateUintSlice(rawValues []string, precision int) (slice interface{}, err error){
	uintSlice := make([]uint, len(rawValues))
	for index, rawValue := range rawValues {
		uintVal, err := strconv.ParseUint(rawValue, 10, precision)
		if err != nil {
			return uintSlice, err
		}
		uintSlice[index] = uint(uintVal)
	}
	return uintSlice, err
}

func generateIntSlice(rawValues []string, precision int) (slice interface{}, err error){
	switch precision {
	case 0:
		intSlice := make([]int, len(rawValues))
		for index, rawValue := range rawValues {
			intVal, err := strconv.ParseInt(rawValue, 10, precision)
			if err != nil {
				return intSlice, err
			}
			intSlice[index] = int(intVal)
		}
		return intSlice, err
	case 8:
		intSlice := make([]int8, len(rawValues))
		for index, rawValue := range rawValues {
			intVal, err := strconv.ParseInt(rawValue, 10, precision)
			if err != nil {
				return intSlice, err
			}
			intSlice[index] = int8(intVal)
		}
		return intSlice, err
	case 16:
		intSlice := make([]int16, len(rawValues))
		for index, rawValue := range rawValues {
			intVal, err := strconv.ParseInt(rawValue, 10, precision)
			if err != nil {
				return intSlice, err
			}
			intSlice[index] = int16(intVal)
		}
		return intSlice, err
	case 32:
		intSlice := make([]int32, len(rawValues))
		for index, rawValue := range rawValues {
			intVal, err := strconv.ParseInt(rawValue, 10, precision)
			if err != nil {
				return intSlice, err
			}
			intSlice[index] = int32(intVal)
		}
		return intSlice, err
	case 64:
		intSlice := make([]int64, len(rawValues))
		for index, rawValue := range rawValues {
			intVal, err := strconv.ParseInt(rawValue, 10, precision)
			if err != nil {
				return intSlice, err
			}
			intSlice[index] = int64(intVal)
		}
		return intSlice, err
	}
	return slice, errors.New("Cannot determine precision of int to generate")
}

func generateFloatSlice(rawValues []string, precision int) (slice interface{}, err error){
	switch precision {
	case 64:
		float64Slice := make([]float64, len(rawValues))
		for index, rawValue := range rawValues {
			float64Val, err := strconv.ParseFloat(rawValue, precision)
			if err != nil {
				return float64Slice, err
			}
			float64Slice[index] = float64(float64Val)
		}
		return float64Slice, err
	case 32:
		float32Slice := make([]float32, len(rawValues))
		for index, rawValue := range rawValues {
			float32Val, err := strconv.ParseFloat(rawValue, precision)
			if err != nil {
				return float32Slice, err
			}
			float32Slice[index] = float32(float32Val)
		}
		return float32Slice, err
	}
	return slice, errors.New("Cannot determine precision of int to generate")

}
