package test

import (
	"testing"

	"5tk.dev/tupi"
)

func TestMapsStringString(t *testing.T) {
	// var m1 = map[string]string{}
	var m = tupi.Validate[map[string]string](map[any]any{"hello": "world"})
	if m.HasErrors() {
		t.Error(m.Errors())
	}
	if h := m.Value()["hello"]; h != "world" {
		t.Errorf("TestMapString got %q, want %q", h, "world")
	}
}

func TestMapsStringInt(t *testing.T) {
	var m = tupi.Validate[map[string]int](map[any]any{"int": "1"})
	if m.HasErrors() {
		t.Error(m.Errors())
	}

	if h := m.Value()["int"]; h != 1 {
		t.Errorf("TestMapString got %q, want %q", h, 1)
	}
}

func TestMapsStringFloat32(t *testing.T) {
	var m = tupi.Validate[map[string]float32](map[any]any{"float": "1.123"})
	if m.HasErrors() {
		t.Error(m.Errors())
	}

	h := m.Value()["float"]
	if h != 1.123 {
		t.Errorf("TestMapString got %v, want %v", h, 1.123)
	}
}

func TestMapsStringBool(t *testing.T) {
	var m = tupi.Validate[map[string]bool](map[any]any{"bool": true})
	if m.HasErrors() {
		t.Error(m.Errors())
	}

	if h, ok := m.Value()["bool"]; !ok {
		t.Errorf("TestMapString got %v, want %v", h, true)
	}
}

func TestMapsStringStruct(t *testing.T) {
	var data = map[any]any{
		"1": &UserTest{Name: "etho"},
		"2": &UserTest{Name: "etho2"},
		"3": &UserTest{Name: "etho3"},
		"4": &UserTest{Name: "etho4"},
		"5": &UserTest{Name: "etho5"},
		"6": &UserTest{Name: "etho6"},
	}
	var m = tupi.Validate[map[string]*UserTest](data)

	if m.HasErrors() {
		t.Error(m.Errors())
	}

	m3 := m.Value()
	if len(m3) != 6 {
		t.Errorf("TestMapString got %v, want %v", m3, data)
	}
	if m3["1"].Name != "etho" {
		t.Errorf("TestMapString['1'] got %v, want %v", m3["1"], data)
	}
	if m3["2"].Name != "etho2" {
		t.Errorf("TestMapString['2'] got %v, want %v", m3["2"], data)
	}
	if m3["3"].Name != "etho3" {
		t.Errorf("TestMapString['3'] got %v, want %v", m3["3"], data)
	}
	if m3["4"].Name != "etho4" {
		t.Errorf("TestMapString['4'] got %v, want %v", m3["4"], data)
	}
	if m3["5"].Name != "etho5" {
		t.Errorf("TestMapString['5'] got %v, want %v", m3["5"], data)
	}
	if m3["6"].Name != "etho6" {
		t.Errorf("TestMapString['6'] got %v, want %v", m3["6"], data)
	}
}
