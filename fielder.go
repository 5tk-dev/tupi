package tupi

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	yaml "gopkg.in/yaml.v3"
)

var (
	_skipTag = map[string]any{
		"realName": nil,
		"strType":  nil,
	}
)

type Fielder[T any] struct {
	*tag
	schema T

	name         string
	rules        map[string]*Rule
	mapTags      map[string]string
	realName     string
	defaultValue any

	reflecKind reflect.Kind
	reflecType reflect.Type

	isMAP,
	isSlice,
	isStruct,
	isPointer bool

	sliceType,
	mapKeyType,
	mapValueType *Fielder[any]

	children      map[string]*Fielder[any] //
	fieldsByIndex map[int]string           //
	superIndex    *int                     // if a field to a struct

	/*
		this permit the use:
			var fielder = &tupi.Fielder[*Foo]{}
			sch := fielder.DecodeFromJson(`{"foo":"bar"}`)
	*/
	parsed bool
}

func (f *Fielder[T]) parseRules() {
	if f.rules == nil {
		f.rules = map[string]*Rule{}
	}
	for rn, v := range f.mapTags {
		if r := GetRule(rn); r != nil {
			nr := &Rule{
				Name:     r.Name,
				Value:    v,
				Message:  r.Message,
				Validate: r.Validate,
			}
			f.rules[rn] = nr
			switch r.Name {
			case "min", "max", "minlen", "maxlen":
				f.required = true
				f.nullable = false
			}
		}
	}
}

func (f *Fielder[T]) checkSchPtr(r reflect.Value) T {
	if f.isPointer && (r.CanAddr() && r.Kind() != reflect.Pointer) {
		return r.Addr().Interface().(T)
	} else if !f.isPointer && r.Kind() == reflect.Pointer {
		return r.Elem().Interface().(T)
	}
	return r.Interface().(T)
}

func (f *Fielder[T]) execRules(sch reflect.Value) (reflect.Value, error) {
	for _, r := range f.rules {
		if !r.Validate(sch, r.Value) {
			return reflect.Value{}, ValidationError{
				Rule:  r,
				Field: f.name,
			}

		}
	}
	return sch, nil
}

func (f *Fielder[T]) decodePrimitive(rv reflect.Value) (reflect.Value, any) {
	if f.reflecKind == reflect.Interface {
		return rv, nil
	}
	sch := f.New()
	if !SetReflectValue(sch, rv) {
		if !f.skipError {
			return reflect.Value{}, RetInvalidType(f)
		}
	}
	if f.isPointer && sch.CanAddr() {
		return sch.Addr(), nil
	}
	return sch, nil
}

func (f *Fielder[T]) decodeSlice(rv reflect.Value) (sch reflect.Value, err any) {

	errs := []any{}
	sliceOf := reflect.TypeOf(f.schema)
	lenSlice := rv.Len()

	if f.isPointer {
		sch = reflect.MakeSlice(sliceOf.Elem(), lenSlice, rv.Cap())
	} else {
		sch = reflect.MakeSlice(sliceOf, lenSlice, rv.Cap())
	}

	for i := range lenSlice {
		var (
			is       = rv.Index(i)
			sf       = f.sliceType
			err      any
			sliceSch reflect.Value
		)

		sliceSch, err = sf.decodeSchema(is.Interface())
		if err != nil {
			errs = append(errs, err)
			continue
		}

		schIndex := sch.Index(i)
		if f.sliceType.isPointer {
			if sliceSch.Kind() != reflect.Ptr && sliceSch.CanAddr() {
				sliceSch = sliceSch.Addr()
			}
		} else {
			if sliceSch.Kind() == reflect.Ptr {
				sliceSch = sliceSch.Elem()
			}
		}
		schIndex.Set(sliceSch)
	}

	if sch.Len() == 0 {
		if f.required {
			errs = append(errs, RetMissing(f))
		}
	}

	if len(errs) == 1 {
		err = errs[0]
	} else if len(errs) > 0 {
		err = errs
	}

	return
}

func (f *Fielder[T]) decodeMap(rv reflect.Value) (sch reflect.Value, err any) {
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		err = map[string]any{f.name: RetInvalidType(f)}
		return
	}
	mt := reflect.TypeOf(f.schema)
	m := reflect.MakeMap(mt)
	for _, key := range rv.MapKeys() {
		mindex := rv.MapIndex(key)

		mkey, _err := f.mapKeyType.decodeSchema(key.Interface())
		if _err != nil {
			err = _err
			return
		}
		mval, _err := f.mapValueType.decodeSchema(mindex.Interface())
		if _err != nil {
			err = _err
			return
		}
		if f.mapValueType.isPointer && mval.CanAddr() {
			mval = mval.Addr()
		}
		if f.mapKeyType.isPointer && mkey.CanAddr() {
			mkey = mkey.Addr()
		}
		m.SetMapIndex(mkey, mval)
	}
	sch = m
	return
}

func (f *Fielder[T]) decodeStruct(v any) (reflect.Value, any) {
	errs := []any{}
	var data map[string]any

	rvData := reflect.ValueOf(v)
	if rvData.Kind() == reflect.Pointer {
		if _rvData := rvData.Elem(); _rvData.Kind() != reflect.Invalid {
			rvData = _rvData
		}
	}
	switch rvData.Kind() {
	case reflect.Map, reflect.Struct:
		_d, err := EncodeStruct(v)
		if err != nil {
			if f.skipError {
				return f.New(), nil
			} else {
				return reflect.Value{}, RetInvalidType(f)
			}
		}
		_data, ok := _d.(map[string]any)
		if !ok {
			if f.skipError {
				return f.New(), nil
			} else {
				return reflect.Value{}, RetInvalidType(f)
			}
		}
		data = _data
	}

	sch := f.New()

	for i := range sch.NumField() {
		schF := sch.Field(i)
		if !schF.CanInterface() {
			continue
		}

		var value any
		fName := f.fieldsByIndex[i]
		fielder, ok := f.children[fName]
		if !ok {
			continue
		}

		if fielder.recursive {
			value = data
		} else {
			if value, ok = data[fName]; !ok {
				if fielder.defaultValue != nil {
					value = fielder.defaultValue
				}
			}
			if value == nil {
				if !fielder.nullable {
					if fielder.required {
						errs = append(errs, RetMissing(fielder))
					}
				}
				continue
			}
		}

		if value == nil && fielder.nullable {
			continue
		}

		var rvF reflect.Value
		var err any

		if fielder.skipValidate {
			SetReflectValue(schF,
				reflect.ValueOf(value))
			continue
		}

		if rvF, err = fielder.decodeSchema(value); err != nil {
			errs = append(errs, err)
			continue
		}

		if rvF.Kind() == reflect.Invalid {
			if fielder.nullable {
				continue
			} else if !fielder.skipError {
				errs = append(errs, map[string]any{fielder.name: RetInvalidType(fielder)})
			}
		} else if !SetReflectValue(schF, rvF) {
			if !fielder.skipError {
				errs = append(errs, RetInvalidType(fielder))
			}
			continue
		}
		if _, ruleErrs := fielder.execRules(schF); ruleErrs != nil {
			errs = append(errs, ruleErrs)
		}
	}
	var err any
	if len(errs) == 1 {
		err = errs[0]
	} else if len(errs) > 0 {
		err = errs
	}

	return sch, err
}

func (f *Fielder[T]) decodeSchema(v any) (reflect.Value, any) {
	if v == "" && f.reflecKind != reflect.String { // if v == a string: (nil or null), v = nil
		v = nil
	}
	var (
		sch reflect.Value
		err any
	)

	if v == nil {
		if f.defaultValue != nil {
			sch, err = f.decodeSchema(f.defaultValue)
		} else if f.required {
			errs := map[string]any{}
			if len(f.children) > 0 {
				for _, c := range f.children {
					if c.required {
						errs[c.name] = RetMissing(c)
					}
				}
				return reflect.Value{}, errs
			} else {
				return reflect.Value{}, map[string]any{
					f.name: RetMissing(f),
				}
			}
		} else if f.nullable {
			sch, err = f.execRules(reflect.ValueOf(nil))
		}
	}
	if err == nil {
		rfVal := reflect.ValueOf(v)
		switch f.reflecKind {
		default:
			sch, err = f.decodePrimitive(rfVal)
		case reflect.Map:
			sch, err = f.decodeMap(rfVal)
		case reflect.Array, reflect.Slice:
			sch, err = f.decodeSlice(rfVal)
		case reflect.Struct:
			sch, err = f.decodeStruct(v)
		}
	}
	if err != nil {
		return reflect.Value{}, err
	}
	return f.execRules(sch)
}

/*
	Decode
*/

func (f *Fielder[T]) Decode(data any) Schema[T] {
	if !f.parsed {
		*f = *parse(*new(T), "validate", map[string]string{})
	}
	sch, err := f.decodeSchema(data)
	s := &schema[T]{}
	if err != nil {
		if e, ok := err.(error); ok {
			s.errors = append(s.errors, e)
			return s
		}
		if str, ok := err.(string); ok {
			s.errors = append(s.errors, errors.New(str))
			return s
		}
		s.errors = append(s.errors, errors.New(fmt.Sprint(err)))
		return s
	}
	s.val = f.checkSchPtr(sch)
	return s
}

func (f *Fielder[T]) DecodeFromYaml(data any) Schema[T] {
	if d, ok := data.(string); ok {
		if (f.reflecKind == reflect.Map) || (f.reflecKind == reflect.Struct) || (f.reflecKind == reflect.Slice) {
			var m any
			err := yaml.Unmarshal([]byte(d), &m)
			if err != nil {
				return &schema[T]{
					errors: []error{RetInvalidType(f)},
				}
			}
			data = m
		}
	}
	return f.Decode(data)
}

func (f *Fielder[T]) DecodeFromJson(data any) Schema[T] {
	if d, ok := data.(string); ok {
		if (f.reflecKind == reflect.Map) || (f.reflecKind == reflect.Struct) || (f.reflecKind == reflect.Slice) {
			var m any
			err := json.Unmarshal([]byte(d), &m)
			if err != nil {
				return &schema[T]{
					errors: []error{RetInvalidType(f)},
				}
			}
			data = m
		}
	}
	return f.Decode(data)
}

/*
	Utils
*/

func (f *Fielder[T]) New() reflect.Value {
	rs := reflect.ValueOf(f.schema)

	if f.isSlice {
		t := reflect.TypeOf(f.sliceType.schema)
		t = reflect.SliceOf(t)
		rs = reflect.MakeSlice(t, 0, 0)
	}

	t := rs.Type()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	v := reflect.New(t).Elem()
	return v
}

func (f *Fielder[T]) ToMap() map[string]any {
	fieldMap := map[string]any{}
	for t, v := range f.mapTags {
		if _, ok := _skipTag[t]; ok {
			continue
		}
		fieldMap[t] = v
	}

	st := f.mapTags["strType"]
	if st == "" {
		st = f.reflecKind.String()
	}

	if st != "struct" {
		fieldMap["type"] = st
	}

	if len(f.children) > 0 {
		for cn, cv := range f.children {
			fieldMap[cn] = cv.ToMap()
		}
	} else if f.isSlice {
		fieldMap["fields"] = f.sliceType.ToMap()
	}

	mapRules := []map[string]any{}
	for _, r := range f.rules {
		mapRules = append(mapRules, r.ToMap())
	}
	fieldMap["rules"] = mapRules

	return fieldMap
}
