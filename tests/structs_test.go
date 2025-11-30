package test

import (
	"testing"

	"5tk.dev/tupi"
)

type TagOmit struct {
	Field string `validate:"-"`
}

// tag "-"
func TestStructTag_omit(t *testing.T) {

	sch := tupi.Validate[*TagOmit](map[string]string{"field": "batata"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	if v := sch.Value(); v.Field != "" {
		t.Errorf("TagOmit.Field: got %q, want %q", v.Field, "")
	}
}

type TagName struct {
	Field string `validate:"name=uuid"`
}

// tag "name"
func TestStructTag_name(t *testing.T) {
	uid := "000-000"
	sch := tupi.Validate[*TagName](map[string]any{"uuid": uid})
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	if v := sch.Value(); v.Field != uid {
		t.Errorf("TagName.Field: got %q, want %q", v.Field, uid)
	}
}

type TagEscape struct {
	Field string `validate:"escape"`
}

// tag "escape"
func TestStructTag_escape(t *testing.T) {
	html := "<h1>Hello</h1>"
	sch := tupi.Validate[*TagEscape](map[string]string{"field": html})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
	if v := sch.Value(); v.Field == html {
		t.Errorf("TagEscape.Field: got %q, want %q", v.Field, html)
	}
}

type TagRequired struct {
	Field string `validate:"required"`
}

// tag "required"
func TestStructTag_required(t *testing.T) {
	sch := tupi.Validate[*TagRequired](map[string]string{"field": ""}) // same result map[string]string{}
	if !sch.HasErrors() {
		v := sch.Value()
		t.Errorf("TagRequired.Field: got %q, want %q", v.Field, "error")
	}
}

type TagNullable struct {
	Field *TagEscape `validate:"nullable"`
}

// tag "nullable"
func TestStructTag_nullable(t *testing.T) {

	sch := tupi.Validate[*TagNullable](map[string]any{})
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}

	v := sch.Value()
	if v.Field != nil {
		t.Errorf("TagNullable.Field: got %v, want \"nil\"", v.Field)
	}
}

type TagSkiperr struct {
	Field int `validate:"skiperr"`
}

// tag "skiperr"
func TestStructTag_skiperr(t *testing.T) {
	sch := tupi.Validate[*TagSkiperr](map[string]any{"field": nil})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
}

type UserTest struct {
	Name string
	Age  int
}

type TagRecursive struct {
	UserTest `validate:"recursive"`
	Role     int
}

// tag "recursive"
func TestStructTag_recursive(t *testing.T) {

	f := tupi.Parse[*TagRecursive]()
	sch := f.Decode(map[string]any{"name": "r2d2", "age": 18, "role": 1})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Role != 1 || v.Age != 18 || v.Name != "r2d2" {
		t.Errorf("TagRecursive.Field: got %v, want %v", v, TagRecursive{
			UserTest: UserTest{Name: "r2d2", Age: 18},
			Role:     1,
		})
	}
}

type TagMinMax struct {
	Num int `validate:"min=1,max=2"`
}

func TestTagStructTag_minMax(t *testing.T) {
	f := tupi.Parse[*TagMinMax]()
	sch := f.Decode(map[string]string{})
	sch0 := f.Decode(map[string]int{"num": 0})
	sch1 := f.Decode(map[string]int{"num": 1})
	sch2 := f.Decode(map[string]int{"num": 2})
	sch3 := f.Decode(map[string]int{"num": 3})
	sch4 := f.Decode(map[string]int{"num": 4})

	if !sch.HasErrors() {
		t.Errorf("sch  -> TagMinMax.Field: got %v, want 'min' error", sch.Value().Num)
	}
	if !sch0.HasErrors() {
		t.Errorf("sch0 -> TagMinMax.Field: got %v, want 'min' error", sch0.Value().Num)
	}
	if sch1.HasErrors() {
		t.Error(sch1.Errors())
	}
	if sch2.HasErrors() {
		t.Error(sch2.Errors())
	}
	if !sch3.HasErrors() {
		t.Errorf("sch3 -> TagMinMax.Field: got %v, want 'max' error", sch3.Value().Num)
	}
	if !sch4.HasErrors() {
		t.Errorf("sch4 -> TagMinMax.Field: got %v, want 'max' error", sch4.Value().Num)
	}
}
