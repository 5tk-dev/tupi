package tupi

import (
	"reflect"
	"strings"
)

/*
usage:

	tupi.Parse(struct{}) => struct{}
	tupi.Parse(&struct{}) => *struct{}
	tupi.Parse(&struct{Field:Value}) => *struct{Field: value} // with default value

	type Schema struct{
		Field string	`validate:"-"`// omit this field
		Field string	`validate:"name"` // string: name of validation	(default realName)
		Field string	`validate:"escape"`	// bool: escape html value (default false)
		Field string	`validate:"required"` // bool (default false)
		Field string	`validate:"nullable"` // bool: if true, allow nil value	(default true)
		Field Struct{}	`validate:"recursive"` // bool: for embbed data (default false)
		Field any		`validate:"skiperr"` // bool: omit on error (default false)
		Field any		`validate:"skip"` // bool: set value without validate	(default false)
		Field int		`validate:"min=18"`	// numbers only (int8, 16..., float32, ...)
		Field int		`validate:"max=65"`	// numbers only (int8, 16..., float32, ...)
		Field string	`validate:"minlength=1"`// if a value can len, is valid. else skip
		Field []string	`validate:"maxlength=100"`// if a value can len, is valid. else skip

		Field string	`validate:"escape,minlen=6,maxlen=64"` // ex. many options
	}
*/
func Parse[T any](schema ...T) *Fielder[T] {
	return ParseWithCustomTag("validate", schema...)
}

func ParseWithCustomTag[T any](tagKey string, schema ...T) *Fielder[T] {
	var sch T
	if len(schema) > 0 {
		sch = schema[0]
	}
	return parse(sch, tagKey, map[string]string{})
}

func parse[T any](schema T, tagKey string, tags map[string]string) *Fielder[T] {
	if _, ok := tags["-"]; ok {
		return nil
	}
	var (
		f  = &Fielder[T]{schema: schema}
		rv = reflect.ValueOf(schema)
		rt = reflect.TypeOf(schema)
	)

	if rt.Kind() == reflect.Pointer {
		f.isPointer = true
		rt = rt.Elem()
		rv = rv.Elem()
	}

	f.realName = tags["realName"]
	f.tag = newTag(tags)
	f.mapTags = tags
	f.parseRules()

	if !rv.IsValid() {
		rv = reflect.New(rt).Elem()
	}

	f.reflecKind = rt.Kind()
	f.reflecType = rt
	f.children = make(map[string]*Fielder[any])

	if f.realName == "" && f.reflecKind != reflect.Interface {
		f.realName = rt.Name()
	}

	switch f.reflecKind {
	case reflect.Struct:
		f.isStruct = true
		f.fieldsByIndex = map[int]string{}
		for i := 0; i < rt.NumField(); i++ {
			fv := rv.Field(i)
			if fv.CanInterface() {
				ft := rt.Field(i)
				childTags := parseTags(ft.Tag.Get(tagKey))
				if _, ok := childTags["-"]; ok {
					continue
				}

				cname := ""
				childTags["realName"] = ft.Name
				if v, ok := childTags["name"]; ok && v != "" {
					cname = v
				} else {
					cname = strings.ToLower(ft.Name)
				}
				childTags["name"] = cname

				var fi any
				if fv.Kind() != reflect.Interface {
					fi = fv.Interface()
				}

				child := parse(fi, tagKey, childTags)
				f.fieldsByIndex[i] = cname
				if child != nil {
					child.superIndex = &i
					f.children[cname] = child
					if v, ok := childTags["recursive"]; ok && strings.ToLower(v) == "true" {
						child.recursive = true
					}
				}
			}
		}
	case reflect.Slice, reflect.Array:
		objIsPrt := false
		f.reflecKind = rt.Kind()
		f.isSlice = true
		rvt := rv.Type().Elem()
		if rvt.Kind() == reflect.Pointer {
			objIsPrt = true
			rvt = rvt.Elem()
		}
		sliceObjet := reflect.New(rvt).Elem()
		f.sliceType = parse[any](sliceObjet.Interface(), tagKey, map[string]string{"realName": ""})
		f.sliceType.isPointer = objIsPrt
	case reflect.Map:
		f.isMAP = true

		keyIsPtr := false
		valIsPtr := false

		mapKey := reflect.New(rt.Key()).Elem()
		if mapKey.Kind() == reflect.Pointer {
			keyIsPtr = true
			mapKey = mapKey.Elem()
		}
		mapValue := reflect.New(rt.Elem()).Elem()

		if mapValue.Kind() == reflect.Pointer {
			valIsPtr = true
		}

		f.mapKeyType = parse(mapKey.Interface(), tagKey, map[string]string{"realName": ""})
		f.mapValueType = parse(mapValue.Interface(), tagKey, map[string]string{"realName": ""})

		f.mapKeyType.isPointer = keyIsPtr
		f.mapValueType.isPointer = valIsPtr
	}
	if rv.IsValid() {
		if rv.CanInterface() && !rv.IsZero() {
			f.defaultValue = schema
			f.required = false
		}
	}
	f.parsed = true
	return f
}
