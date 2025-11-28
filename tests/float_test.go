package test

import (
	"testing"

	"5tk.dev/tupi"
)

var (
	fielderFloat32 = tupi.ParseSchema[float32]()
	fielderFloat64 = tupi.ParseSchema[float64]()
)

func TestFloatsWithFloats(t *testing.T) {
	data64 := 1.9223372036854775807
	sch32 := fielderFloat32.Decode(data64)
	if sch32.HasErrors() {
		t.Error(sch32.Errors())
	}
	v32 := sch32.Value()
	if v32 != float32(data64) {
		t.Errorf("got %v, want %v", v32, data64)
	}

	sch64 := fielderFloat64.Decode(data64)
	if sch64.HasErrors() {
		t.Error(sch64.Errors())
	}
	v64 := sch64.Value()
	if v64 != 1.9223372036854775807 {
		t.Errorf("got %v, want %v", v64, data64)
	}

}

// floats.Decode(int) -> !Schema.hasError()
func TestFloatsWithInts(t *testing.T) {
	data64 := 2147483647
	sch32 := fielderFloat32.Decode(data64)
	if sch32.HasErrors() {
		t.Error(sch32.Errors())
	}
	v32 := sch32.Value()
	if v32 != float32(data64) {
		t.Errorf("got %v, want %v", v32, float32(data64))
	}

	sch64 := fielderFloat64.Decode(data64)
	if sch64.HasErrors() {
		t.Error(sch64.Errors())
	}
	v64 := sch64.Value()
	if v64 != float64(data64) {
		t.Errorf("got %v, want %v", v64, (float64(data64)))
	}
}

// floats.Decode("int") -> !Schema.hasError()
func TestFloatsWithStringInts(t *testing.T) {

	data64 := "214748365"
	sch32 := fielderFloat32.Decode(data64)
	if sch32.HasErrors() {
		t.Error(sch32.Errors())
	}
	v32 := sch32.Value()
	if v32 != float32(214748365) {
		t.Errorf("got %v, want %v", v32, data64)
	}

	sch64 := fielderFloat64.Decode(data64)
	if sch64.HasErrors() {
		t.Error(sch64.Errors())
	}
	v64 := sch64.Value()
	if v64 != float64(214748365) {
		t.Errorf("got %v, want %v", v64, data64)
	}
}

// floats.Decode("string") -> Schema.hasError()
func TestFloatsWithStringValues(t *testing.T) {
	sch := fielderFloat32.Decode("\t\t")
	if !sch.HasErrors() {
		t.Errorf("got %v, want %q", sch.Value(), "error")
	}

	sch64 := fielderFloat64.Decode("\r\n")
	if !sch64.HasErrors() {
		t.Errorf("got %v, want %q", sch64.Value(), "error")
	}
}

// floats.Decode(false) -> Schema.hasError()
func TestFloatsWithBooleans(t *testing.T) {
	sch32 := fielderFloat32.Decode(false)
	if !sch32.HasErrors() {
		t.Errorf("got %v, want %q", sch32.Value(), "error")
	}

	sch64 := fielderFloat64.Decode(false)
	if !sch64.HasErrors() {
		t.Errorf("got %v, want %q", sch64.Value(), "error")
	}
}

// floats.Decode(true) -> Schema.hasError()
func TestFloatsWithBooleans1(t *testing.T) {
	sch32 := fielderFloat32.Decode(true)
	if !sch32.HasErrors() {
		t.Errorf("got %v, want %q", sch32.Value(), "error")
	}
	sch64 := fielderFloat64.Decode(true)
	if !sch64.HasErrors() {
		t.Errorf("got %v, want %q", sch64.Value(), "error")
	}
}

// floats.Decode(struct) -> Schema.hasError()
func TestFloatsWithStructs(t *testing.T) {
	s := struct{}{}
	sch32 := fielderFloat32.Decode(s)
	if !sch32.HasErrors() {
		t.Errorf("got %v, want %q", sch32.Value(), "error")
	}

	sch64 := fielderFloat64.Decode(s)
	if !sch64.HasErrors() {
		t.Errorf("got %v, want %q", sch64.Value(), "error")
	}
}

// floats.Decode(map) -> Schema.hasError()
func TestFloatsWithMaps(t *testing.T) {
	m := map[string]any{}
	sch32 := fielderFloat32.Decode(m)
	if !sch32.HasErrors() {
		t.Errorf("got %v, want %q", sch32.Value(), "error")
	}
	sch64 := fielderFloat32.Decode(m)
	if !sch64.HasErrors() {
		t.Errorf("got %v, want %q", sch64.Value(), "error")
	}
}

// floats.Decode(func) -> Schema.hasError()
func TestFloatsWithFuncs(t *testing.T) {
	f := func() {}
	sch32 := fielderFloat32.Decode(f)
	if !sch32.HasErrors() {
		t.Errorf("got %v, want %q", sch32.Value(), "error")
	}
	sch := fielderFloat64.Decode(f)
	if !sch.HasErrors() {
		t.Errorf("got %v, want %q", sch.Value(), "error")
	}
}
