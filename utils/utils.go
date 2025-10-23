package utils

import (
	"fmt"
	"reflect"
	"sort"
	"time"

	"taskmanager/model"
)

const DateLayout = "02-01-2006"

func SortTasks(tasks []model.Task, fieldname string) ([]model.Task, error) {
	taskType := reflect.TypeOf(model.Task{})

	field, ok := taskType.FieldByName(fieldname)
	if !ok {
		return nil, fmt.Errorf("No such field %s in struct", fieldname)
	}

	sort.Slice(tasks, func(i, j int) bool {
		valI := reflect.ValueOf(tasks[i]).FieldByName(fieldname)
		valJ := reflect.ValueOf(tasks[j]).FieldByName(fieldname)

		switch field.Type.Kind() {
		case reflect.String:
			return valI.String() < valJ.String()
		case reflect.Int:
			return valI.Int() < valJ.Int()
		case reflect.Struct:
			if field.Type == reflect.TypeOf(model.Status("")) {
				return valI.String() < valJ.String()
			} else if field.Type == reflect.TypeOf(time.Time{}) {
				return valI.Interface().(time.Time).Before(valJ.Interface().(time.Time))
			}
			return false
		default:
			return false
		}
	})

	return tasks, nil
}

func SetField(obj any, fieldname string, value any) error {
	v := reflect.ValueOf(obj).Elem()
	field := v.FieldByName(fieldname)

	if !field.IsValid() {
		return fmt.Errorf("No such field: %s in object", fieldname)
	}
	if !field.CanSet() {
		return fmt.Errorf("Cannot set field: %s", fieldname)
	}

	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return fmt.Errorf("Provided value type did not match object field type")
	}
	field.Set(val)

	return nil
}
