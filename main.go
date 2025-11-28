package tupi

// func init() { initRegistry() }

// use 'validate' in tags
//
//	type User struct{
//		Name string `validate:"minlen=4"`
//		Age int `validate:"min=18"`
//	}
//	sch := tupi.Validate(&User{},map[string]any{})
//	if sch.HasErrors(){
//		err := sch.Errors()
//		....
//	} else {
//		user := sch.Value().(*User)
//	}
func Validate[T any](data any) Schema[T] {
	return ParseSchemaWithTag[T]("validate").Decode(data)
}
