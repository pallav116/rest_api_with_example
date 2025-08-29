package common

import (
	"encoding/json"
	"errors"
	"reflect"
)

var ErrNotStruct = errors.New("data must be struct or pointer to struct")

// TODO: must write struct to struct mapper instead of this function
func TypeConverter[T any](data any) (T, error) {
	var result T
	dataJson, err := json.Marshal(&data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dataJson, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Convert maps fields from data (struct or pointer to struct) to a new instance of type T, supporting nested structs.
func Convert[T any](data any) (T, error) {
	var result T
	srcVal := reflect.ValueOf(data)
	if srcVal.Kind() == reflect.Pointer {
		srcVal = srcVal.Elem()
	}

	if srcVal.Kind() != reflect.Struct && !(srcVal.Kind() == reflect.Ptr && srcVal.Elem().Kind() == reflect.Struct) {
		return result, ErrNotStruct
	}

	dstVal := reflect.ValueOf(&result).Elem()
	if err := deepCopy(dstVal, srcVal); err != nil {
		return result, err
	}
	return result, nil
}

// deepCopy recursively copies matching fields from src to dst, including nested structs.
func deepCopy(dst, src reflect.Value) error {
	for i := 0; i < dst.NumField(); i++ {
		dstField := dst.Field(i)
		dstTypeField := dst.Type().Field(i)
		srcField := src.FieldByName(dstTypeField.Name)
		if !srcField.IsValid() {
			continue
		}
		// If field is a struct, recurse
		if dstField.Kind() == reflect.Struct && srcField.Kind() == reflect.Struct {
			if err := deepCopy(dstField, srcField); err != nil {
				return err
			}
		} else if srcField.Type().AssignableTo(dstField.Type()) && dstField.CanSet() {
			dstField.Set(srcField)
		}
	}
	return nil
}
