package pgtype

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

// underlyingNumberType gets the underlying type that can be converted to Int2, Int4, Int8, Float4, or Float8
func underlyingNumberType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.Int:
		convVal := int(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int8:
		convVal := int8(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int16:
		convVal := int16(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int32:
		convVal := int32(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Int64:
		convVal := int64(refVal.Int())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint:
		convVal := uint(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint8:
		convVal := uint8(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint16:
		convVal := uint16(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint32:
		convVal := uint32(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Uint64:
		convVal := uint64(refVal.Uint())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Float32:
		convVal := float32(refVal.Float())
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.Float64:
		convVal := refVal.Float()
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	case reflect.String:
		convVal := refVal.String()
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	}

	return nil, false
}

// underlyingBoolType gets the underlying type that can be converted to Bool
func underlyingBoolType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.Bool:
		convVal := refVal.Bool()
		return convVal, reflect.TypeOf(convVal) != refVal.Type()
	}

	return nil, false
}

// underlyingPtrType dereferences a pointer
func underlyingPtrType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	}

	return nil, false
}

// underlyingTimeType gets the underlying type that can be converted to time.Time
func underlyingTimeType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return time.Time{}, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	}

	timeType := reflect.TypeOf(time.Time{})
	if refVal.Type().ConvertibleTo(timeType) {
		return refVal.Convert(timeType).Interface(), true
	}

	return time.Time{}, false
}

// underlyingSliceType gets the underlying slice type
func underlyingSliceType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	switch refVal.Kind() {
	case reflect.Ptr:
		if refVal.IsNil() {
			return nil, false
		}
		convVal := refVal.Elem().Interface()
		return convVal, true
	case reflect.Slice:
		baseSliceType := reflect.SliceOf(refVal.Type().Elem())
		if refVal.Type().ConvertibleTo(baseSliceType) {
			convVal := refVal.Convert(baseSliceType)
			return convVal.Interface(), reflect.TypeOf(convVal.Interface()) != refVal.Type()
		}
	}

	return nil, false
}

func underlyingPtrSliceType(val interface{}) (interface{}, bool) {
	refVal := reflect.ValueOf(val)

	if refVal.Kind() != reflect.Ptr {
		return nil, false
	}
	if refVal.IsNil() {
		return nil, false
	}

	sliceVal := refVal.Elem().Interface()
	baseSliceType := reflect.SliceOf(reflect.TypeOf(sliceVal).Elem())
	ptrBaseSliceType := reflect.PtrTo(baseSliceType)

	if refVal.Type().ConvertibleTo(ptrBaseSliceType) {
		convVal := refVal.Convert(ptrBaseSliceType)
		return convVal.Interface(), reflect.TypeOf(convVal.Interface()) != refVal.Type()
	}

	return nil, false
}

func int64AssignTo(srcVal int64, srcStatus Status, dst interface{}) error {
	if srcStatus == Present {
		switch v := dst.(type) {
		case *int:
			if srcVal < int64(minInt) {
				return fmt.Errorf("%d is less than minimum value for int", srcVal)
			} else if srcVal > int64(maxInt) {
				return fmt.Errorf("%d is greater than maximum value for int", srcVal)
			}
			*v = int(srcVal)
		case *int8:
			if srcVal < math.MinInt8 {
				return fmt.Errorf("%d is less than minimum value for int8", srcVal)
			} else if srcVal > math.MaxInt8 {
				return fmt.Errorf("%d is greater than maximum value for int8", srcVal)
			}
			*v = int8(srcVal)
		case *int16:
			if srcVal < math.MinInt16 {
				return fmt.Errorf("%d is less than minimum value for int16", srcVal)
			} else if srcVal > math.MaxInt16 {
				return fmt.Errorf("%d is greater than maximum value for int16", srcVal)
			}
			*v = int16(srcVal)
		case *int32:
			if srcVal < math.MinInt32 {
				return fmt.Errorf("%d is less than minimum value for int32", srcVal)
			} else if srcVal > math.MaxInt32 {
				return fmt.Errorf("%d is greater than maximum value for int32", srcVal)
			}
			*v = int32(srcVal)
		case *int64:
			if srcVal < math.MinInt64 {
				return fmt.Errorf("%d is less than minimum value for int64", srcVal)
			} else if srcVal > math.MaxInt64 {
				return fmt.Errorf("%d is greater than maximum value for int64", srcVal)
			}
			*v = int64(srcVal)
		case *uint:
			if srcVal < 0 {
				return fmt.Errorf("%d is less than zero for uint", srcVal)
			} else if uint64(srcVal) > uint64(maxUint) {
				return fmt.Errorf("%d is greater than maximum value for uint", srcVal)
			}
			*v = uint(srcVal)
		case *uint8:
			if srcVal < 0 {
				return fmt.Errorf("%d is less than zero for uint8", srcVal)
			} else if srcVal > math.MaxUint8 {
				return fmt.Errorf("%d is greater than maximum value for uint8", srcVal)
			}
			*v = uint8(srcVal)
		case *uint16:
			if srcVal < 0 {
				return fmt.Errorf("%d is less than zero for uint32", srcVal)
			} else if srcVal > math.MaxUint16 {
				return fmt.Errorf("%d is greater than maximum value for uint16", srcVal)
			}
			*v = uint16(srcVal)
		case *uint32:
			if srcVal < 0 {
				return fmt.Errorf("%d is less than zero for uint32", srcVal)
			} else if srcVal > math.MaxUint32 {
				return fmt.Errorf("%d is greater than maximum value for uint32", srcVal)
			}
			*v = uint32(srcVal)
		case *uint64:
			if srcVal < 0 {
				return fmt.Errorf("%d is less than zero for uint64", srcVal)
			}
			*v = uint64(srcVal)
		default:
			if v := reflect.ValueOf(dst); v.Kind() == reflect.Ptr {
				el := v.Elem()
				switch el.Kind() {
				// if dst is a pointer to pointer, strip the pointer and try again
				case reflect.Ptr:
					if el.IsNil() {
						// allocate destination
						el.Set(reflect.New(el.Type().Elem()))
					}
					return int64AssignTo(srcVal, srcStatus, el.Interface())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					if el.OverflowInt(int64(srcVal)) {
						return fmt.Errorf("cannot put %d into %T", srcVal, dst)
					}
					el.SetInt(int64(srcVal))
					return nil
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					if srcVal < 0 {
						return fmt.Errorf("%d is less than zero for %T", srcVal, dst)
					}
					if el.OverflowUint(uint64(srcVal)) {
						return fmt.Errorf("cannot put %d into %T", srcVal, dst)
					}
					el.SetUint(uint64(srcVal))
					return nil
				}
			}
			return fmt.Errorf("cannot assign %v into %T", srcVal, dst)
		}
		return nil
	}

	// if dst is a pointer to pointer and srcStatus is not Present, nil it out
	if v := reflect.ValueOf(dst); v.Kind() == reflect.Ptr {
		el := v.Elem()
		if el.Kind() == reflect.Ptr {
			el.Set(reflect.Zero(el.Type()))
			return nil
		}
	}

	return fmt.Errorf("cannot assign %v %v into %T", srcVal, srcStatus, dst)
}

func float64AssignTo(srcVal float64, srcStatus Status, dst interface{}) error {
	if srcStatus == Present {
		switch v := dst.(type) {
		case *float32:
			*v = float32(srcVal)
		case *float64:
			*v = srcVal
		default:
			if v := reflect.ValueOf(dst); v.Kind() == reflect.Ptr {
				el := v.Elem()
				switch el.Kind() {
				// if dst is a pointer to pointer, strip the pointer and try again
				case reflect.Ptr:
					if el.IsNil() {
						// allocate destination
						el.Set(reflect.New(el.Type().Elem()))
					}
					return float64AssignTo(srcVal, srcStatus, el.Interface())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					i64 := int64(srcVal)
					if float64(i64) == srcVal {
						return int64AssignTo(i64, srcStatus, dst)
					}
				}
			}
			return fmt.Errorf("cannot assign %v into %T", srcVal, dst)
		}
		return nil
	}

	// if dst is a pointer to pointer and srcStatus is not Present, nil it out
	if v := reflect.ValueOf(dst); v.Kind() == reflect.Ptr {
		el := v.Elem()
		if el.Kind() == reflect.Ptr {
			el.Set(reflect.Zero(el.Type()))
			return nil
		}
	}

	return fmt.Errorf("cannot assign %v %v into %T", srcVal, srcStatus, dst)
}
