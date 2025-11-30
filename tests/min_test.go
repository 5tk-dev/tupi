package test

import (
	"testing"

	"5tk.dev/tupi"
)

type TagMinInt8 struct {
	Value int8 `validate:"min=0"`
}

func TestMinInt8_ok(t *testing.T) {
	sch := tupi.Validate[*TagMinInt8](map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}

	v := sch.Value()
	if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
		return
	}
}

func TestMinInt8_minor(t *testing.T) {
	sch := tupi.Validate[*TagMinInt8](map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMinInt16 struct {
	Value int16 `validate:"min=0"`
}

func TestMinInt16_ok(t *testing.T) {
	sch := tupi.Validate[*TagMinInt16](map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}

	v := sch.Value()
	if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
		return
	}
}

func TestMinInt16_minor(t *testing.T) {
	sch := tupi.Validate[*TagMinInt16](map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMinInt32 struct {
	Value int32 `validate:"min=0"`
}

func TestMinInt32_ok(t *testing.T) {
	sch := tupi.Validate[*TagMinInt32](map[string]any{"value": "1"})
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
		return
	}
}

func TestMinInt32_minor(t *testing.T) {
	sch := tupi.Validate[*TagMinInt32](map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMinInt64 struct {
	Value int64 `validate:"min=0"`
}

func TestMinInt64_ok(t *testing.T) {
	sch := tupi.Validate[*TagMinInt64](map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != 1 {
		t.Errorf("got: %v, want: 1", v.Value)
		return
	}
}

func TestMinInt64_minor(t *testing.T) {
	sch := tupi.Validate[*TagMinInt64](map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMinInt struct {
	Value int `validate:"min=0"`
}

func TestMinInt_ok(t *testing.T) {
	sch := tupi.Validate[*TagMinInt](map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}

	v := sch.Value()
	if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
		return
	}
}

func TestMinInt_minor(t *testing.T) {
	sch := tupi.Validate[*TagMinInt](map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

/*
	FLOATS
*/

type TagMinFloat32 struct {
	Value float32 `validate:"min=1.32"`
}

func TestMinFloat32_ok(t *testing.T) {
	sch := tupi.Validate[*TagMinFloat32](map[string]any{"value": "1.32"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != 1.32 {
		t.Errorf("got: %v, want: 1.32", sch.Value())
		return
	}
}

func TestMinFloat32_minor(t *testing.T) {
	sch := tupi.Validate[*TagMinFloat32](map[string]any{"value": 1.31})

	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMinFloat64 struct {
	Value float64 `validate:"min=0"`
}

func TestMinFloat64_ok(t *testing.T) {
	sch := tupi.Validate[*TagMinFloat64](map[string]any{"value": "1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != 1 {
		t.Errorf("got: %v, want: 1", sch.Value())
		return
	}
}

func TestMinFloat64_minor(t *testing.T) {
	sch := tupi.Validate[*TagMinFloat64](map[string]any{"value": "-1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}
