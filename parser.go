package tupi

import (
	"reflect"
	"strings"
)

/*
usage:

	tupi.ParseSchema(struct{}) => struct{}
	tupi.ParseSchema(&struct{}) => *struct{}
	tupi.ParseSchema(&struct{Field:Value}) => *struct{Field: value} // with default value

	type Schema struct{
		Field `tupi:"-"`				// omit this field
		Field `tupi:"name"`				// string: name of validation		(default realName)
		Field `tupi:"walk"`				// bool: deep validation			(default true)
		Field `tupi:"escape"`			// bool: escape html value			(default false)
		Field `tupi:"required"`			// bool:		...			 		(default false)
		Field `tupi:"nullable"`			// bool: if true, allow nil value	(default true)
		Field `tupi:"recursive"`		// bool: for embbed data 			(default false)
		Field `tupi:"skiperr"`			// bool: omit on error				(default false)
		Field `tupi:"skip"`				// bool: set value without validate	(default false)
		Field `tupi:"min=18"`			// numbers only (int8, 16..., float32, ...)
		Field `tupi:"max=65"`			// numbers only (int8, 16..., float32, ...)
		Field `tupi:"minlength=1"`		// if a value can len, is valid. else skip
		Field `tupi:"maxlength=100"`	// if a value can len, is valid. else skip
	}
*/
func ParseSchema(schema any) *Fielder {
	return ParseSchemaWithTag("tupi", schema)
}

func ParseSchemaWithTag(tagKey string, schema any) *Fielder {
	tags := map[string]string{}
	if rn := reflect.TypeOf(schema).Name(); rn != "" {
		tags["realName"] = rn
	}
	return parseSchema(schema, tagKey, tags)
}

func parseSchema(schema any, tagKey string, tags map[string]string) *Fielder {
	if _, ok := tags["-"]; ok {
		return nil
	}
	var (
		f  = &Fielder{Schema: schema}
		rv = reflect.ValueOf(schema)
		rt reflect.Type
	)
	f.RealName = tags["realName"]

	f.parseTags(tags)
	f.parseRules()

	if schema != nil {
		rt = rv.Type()
	} else {
		rt = reflect.TypeOf(nil)
		f.Type = reflect.Interface
	}

	if rv.Kind() == reflect.Pointer {
		f.IsPointer = true
		rv = rv.Elem()
		rt = rt.Elem()
	}

	if schema != nil {
		f.Type = rv.Kind()
		f.Children = make(map[string]*Fielder)
	}

	if f.RealName == "" && f.Type != reflect.Interface {
		f.RealName = rt.Name()
	}

	switch f.Type {
	case reflect.Struct:
		f.IsStruct = true
		f.FieldsByIndex = map[int]string{}
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

				child := parseSchema(fi, tagKey, childTags)
				f.FieldsByIndex[i] = cname
				if child != nil {
					child.SuperIndex = &i
					f.Children[cname] = child
					if v, ok := childTags["heritage"]; ok && strings.ToLower(v) == "true" {
						child.Recurcive = true
					}
				}
			}
		}
	case reflect.Slice, reflect.Array:
		objIsPrt := false
		f.Type = rt.Kind()
		f.IsSlice = true
		rvt := rv.Type().Elem()
		if rvt.Kind() == reflect.Pointer {
			objIsPrt = true
			rvt = rvt.Elem()
		}
		sliceObjet := reflect.New(rvt).Elem()
		f.SliceType = parseSchema(sliceObjet.Interface(), tagKey, map[string]string{"realName": ""})
		f.SliceType.IsPointer = objIsPrt
	case reflect.Map:
		f.IsMAP = true

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
			// mapValue = mapValue.Elem()
		}

		f.MapKeyType = parseSchema(mapKey.Interface(), tagKey, map[string]string{"realName": ""})
		f.MapValueType = parseSchema(mapValue.Interface(), tagKey, map[string]string{"realName": ""})

		f.MapKeyType.IsPointer = keyIsPtr
		f.MapValueType.IsPointer = valIsPtr
	}
	if rv.IsValid() {
		if rv.CanInterface() && !rv.IsZero() {
			f.Default = schema
		}
	}

	return f
}
