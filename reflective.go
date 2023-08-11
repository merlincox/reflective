package reflective

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
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
		return fmt.Errorf("the argument to FillRandomly to must be a pointer")
	}

	return fillRandomly(val.Elem())
}

func FillRandomlyByValue(val reflect.Value) error {
	if !val.CanSet() {
		return fmt.Errorf("the argument to FillRandomlyByValue must be able to be set")
	}

	return fillRandomly(val)
}

func xfillRandomly(val reflect.Value) error {
	if !val.CanSet() {
		return nil
	}

	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		return fillRandomly(val.Elem())

	case reflect.Bool:
		randBool := (rand.Int() % 2) == 0
		val.SetBool(randBool)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		min64, max64 := int64(math.MinInt64), int64(math.MaxInt64)
		val.SetInt(randInt64(min64, max64))

	case reflect.Uint16, reflect.Uint32, reflect.Uint64:
		randUint := uint64(rand.Uint32())
		val.SetUint(randUint)

	case reflect.Float32, reflect.Float64:
		randFloat := float64(rand.Float32())
		val.SetFloat(randFloat)

	case reflect.String:
		randStringVal := randString()
		val.SetString(randStringVal)

	case reflect.Slice:
		elementType := val.Type().Elem()
		size := 1 + rand.Intn(16)
		sliceVal := reflect.MakeSlice(reflect.SliceOf(elementType), 0, size)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(elementType))
			if err := fillRandomly(newElement); err != nil {
				return err
			}
			sliceVal = reflect.Append(sliceVal, newElement)
		}
		val.Set(sliceVal)

	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if err := fillRandomly(val.Index(i)); err != nil {
				return err
			}
		}

	case reflect.Map:
		mapType := val.Type()
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
		val.Set(mapVal)

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			if err := fillRandomly(val.Field(i)); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported kind: %s", val.Kind().String())
	}
	return nil
}

func MapUtoi(n uint64) int64 {
	return int64(n - 1<<63)
}

func MapItou(n int64) uint64 {
	if n >= 0 {
		return uint64(n) + 1<<63
	}
	return uint64(n + math.MaxInt64 + 1)
}

func irandInt64[T constraints.Signed](n1, n2 T) int64 {
	return randInt64(int64(n1), int64(n2))
}

// Inclusive pseudorandom number. Result is equal to or between n1 and n2.
func randInt64(n1, n2 int64) int64 {
	if n1 == n2 {
		return n2
	}

	if n2 < n1 {
		n2, n1 = n1, n2
	}

	min, max := MapItou(n1), MapItou(n2)

	d := max - min

	if d == math.MaxUint64 {
		return MapUtoi(rand.Uint64())
	}

	if d <= math.MaxInt64 {
		return MapUtoi(min + randInt64i(int64(d)))
	}

	f := rand.Float64()

	if f <= float64(float64(math.MaxInt64)/float64(d)) {
		return MapUtoi(min + uint64(rand.Int63()))
	}

	return MapUtoi(min + math.MaxInt64 + uint64(rand.Int63n(int64(d-math.MaxInt64))))
}

// n inclusive
func randInt64i(max int64) uint64 {
	if max == math.MaxInt64 {
		return uint64(rand.Int63())
	}
	return uint64(rand.Int63n(max + 1))
}
