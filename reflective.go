package reflective

import (
	"fmt"
	"math/rand"
	"reflect"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString() string {
	b := make([]rune, 16)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func FillRandomly(a any) error {
	val := reflect.ValueOf(a)
	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("the argument FillRandomly to must be a pointer")
	}

	return fillRandomly(val.Elem())
}

func fillRandomly(receiver reflect.Value) error {
	if !receiver.CanSet() {
		return nil
	}

	switch receiver.Kind() {
	case reflect.Ptr:
		if receiver.IsNil() {
			receiver.Set(reflect.New(receiver.Type().Elem()))
		}
		return fillRandomly(receiver.Elem())

	case reflect.Bool:
		randBool := (rand.Int() % 2) == 0
		receiver.SetBool(randBool)

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		randInt := int64(rand.Intn(204800))
		receiver.SetInt(randInt)

	case reflect.Uint16, reflect.Uint32, reflect.Uint64:
		randUint := uint64(rand.Uint32())
		receiver.SetUint(randUint)

	case reflect.Float32, reflect.Float64:
		randFloat := float64(rand.Float32())
		receiver.SetFloat(randFloat)

	case reflect.String:
		randStringVal := randString()
		receiver.SetString(randStringVal)

	case reflect.Slice:
		elementType := receiver.Type().Elem()
		size := 1 + rand.Intn(16)
		sliceVal := reflect.MakeSlice(reflect.SliceOf(elementType), 0, size)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(elementType))
			if err := fillRandomly(newElement); err != nil {
				return err
			}
			sliceVal = reflect.Append(sliceVal, newElement)
		}
		receiver.Set(sliceVal)

	case reflect.Array:
		for i := 0; i < receiver.Len(); i++ {
			if err := fillRandomly(receiver.Index(i)); err != nil {
				return err
			}
		}

	case reflect.Map:
		mapType := receiver.Type()
		mapVal := reflect.MakeMap(mapType)
		size := 1 + rand.Intn(16)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(mapType.Elem()))
			if err := fillRandomly(newElement); err != nil {
				return err
			}
			newKey := reflect.Indirect(reflect.New(mapType.Key()))
			if err := fillRandomly(newKey); err != nil {
				return err
			}
			mapVal.SetMapIndex(newKey, newElement)
		}
		receiver.Set(mapVal)

	case reflect.Struct:
		for i := 0; i < receiver.NumField(); i++ {
			if err := fillRandomly(receiver.Field(i)); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported kind: %s", receiver.Kind().String())
	}
	return nil
}
