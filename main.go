package tupi

// use 'validate' in tags
//
//	type User struct{
//		Name string `validate:"minlen=4"`
//		Age int `validate:"min=18"`
//	}
//	sch := tupi.Validate[*User](map[string]any{})
//	if sch.HasErrors(){
//		err := sch.Errors()
//		....
//	}
//	var user *User = sch.Value()
func Validate[T any](data any) Schema[T] {
	return ParseWithCustomTag[T]("validate").Decode(data)
}
