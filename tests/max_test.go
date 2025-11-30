package test

import (
	"testing"

	"5tk.dev/tupi"
)

type TagMaxInt8 struct {
	Value int8 `validate:"max=0"`
}

func TestMaxInt8_ok(t *testing.T) {

	sch := tupi.Validate[*TagMaxInt8](map[string]any{"value": "-1"})
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != -1 {
		t.Errorf("got: %v, want: -1", v.Value)
		return
	}
}

func TestMaxInt8_max(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt8](map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMaxInt16 struct {
	Value int16 `validate:"max=0"`
}

func TestMaxInt16_ok(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt16](map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != -1 {
		t.Errorf("got: %v, want: -1", v.Value)
		return
	}
}

func TestMaxInt16_max(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt16](map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMaxInt32 struct {
	Value int32 `validate:"max=0"`
}

func TestMaxInt32_ok(t *testing.T) {
	sch := tupi.Validate[TagMaxInt32](map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != -1 {
		t.Errorf("got: %v, want: -1", v.Value)
		return
	}
}

func TestMaxInt32_max(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt32](map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMaxInt64 struct {
	Value int64 `validate:"max=0"`
}

func TestMaxInt64_ok(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt64](map[string]any{"value": "-1"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	v := sch.Value()
	if v.Value != -1 {
		t.Errorf("got: %v, want: -1", v.Value)
		return
	}
}

func TestMaxInt64_max(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt64](map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMaxInt struct {
	Value int `validate:"max=0"`
}

func TestMaxInt_ok(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt](map[string]any{"value": "-1"})
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}

	v := sch.Value()
	if v.Value != -1 {
		t.Errorf("got: %v, want: -1", v.Value)
		return
	}
}

func TestMaxInt_max(t *testing.T) {
	sch := tupi.Validate[*TagMaxInt](map[string]any{"value": "1"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	v := sch.Value()
	if v != nil {
		t.Errorf("got: %v, want: nil", v)
		return
	}
}

/*
	FLOATS
*/

type TagMaxFloat32 struct {
	Value float32 `validate:"max=0"`
}

func TestMaxFloat32_ok(t *testing.T) {
	sch := tupi.Validate[*TagMaxFloat32](map[string]any{"value": "-1.32"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}

	v := sch.Value()
	if v.Value != -1.32 {
		t.Errorf("got: %v, want: -1.32", v.Value)
		return
	}
}

func TestMaxFloat32_max(t *testing.T) {
	sch := tupi.Validate[*TagMaxFloat32](map[string]any{"value": "1.32"})

	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}

	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}

type TagMaxFloat64 struct {
	Value float64 `validate:"max=0"`
}

func TestMaxFloat64_ok(t *testing.T) {
	sch := tupi.Validate[*TagMaxFloat64](map[string]any{"value": "-1.64"})

	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}

	v := sch.Value()
	if v.Value != -1.64 {
		t.Errorf("got: %v, want: -1.64", v.Value)
		return
	}
}

func TestMaxFloat64_max(t *testing.T) {
	sch := tupi.Validate[*TagMaxFloat64](map[string]any{"value": "1.64"})
	if !sch.HasErrors() {
		t.Errorf("got: %v, want: error", sch.Value())
		return
	}
	if sch.Value() != nil {
		t.Errorf("got: %v, want: nil", sch.Value())
		return
	}
}
