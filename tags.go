package tupi

type tag struct {

	/*
		is... required or no... (default false)
		type A struct {
			Foo string `tupi:"required"`
		}*/
	required bool

	/*
		allow zero value (default true)
		type B struct {
			A `tupi:"nullable"`
		}
	*/
	nullable bool

	/*
		recursive validation (default false)
		type Model struct {
			UUID string
			CreatedAt time.Time
		}
		type User struct {
			Model
			Name string
		}
		tupi.Validate[User](map[string]any{
			"uuid":"a.b.c.d"
			"name":"janio quadros"
		})
	*/
	recursive bool

	/*
		basicamente: se der erro ele pula pro proximo
	*/
	skipError bool

	/* ... */
	omitEmpty bool

	/*
		se tiver um valor ele setta, senão, vida q segue
	*/
	skipValidate bool
}

func newTag(tags map[string]string) *tag {
	t := &tag{}
	for k, v := range tags {
		switch k {
		case "required":
			t.required = v == "true"
		case "nullable":
			t.nullable = v == "true"
		case "recursive":
			t.recursive = v == "true"
		case "skiperror":
			t.skipError = v == "true"
		case "omitempty":
			t.omitEmpty = v == "true"
		case "skipvalidate":
			t.skipValidate = v == "true"
		}
	}
	return t
}
