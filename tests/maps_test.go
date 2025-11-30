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
		return
	}
	if h := m.Value()["hello"]; h != "world" {
		t.Errorf("TestMapString got %q, want %q", h, "world")
		return
	}
}

func TestMapsStringInt(t *testing.T) {
	var m = tupi.Validate[map[string]int](map[any]any{"int": "1"})
	if m.HasErrors() {
		t.Error(m.Errors())
		return
	}

	if h := m.Value()["int"]; h != 1 {
		t.Errorf("TestMapString got %q, want %q", h, 1)
		return
	}
}

func TestMapsStringFloat32(t *testing.T) {
	var m = tupi.Validate[map[string]float32](map[any]any{"float": "1.123"})
	if m.HasErrors() {
		t.Error(m.Errors())
		return
	}

	h := m.Value()["float"]
	if h != 1.123 {
		t.Errorf("TestMapString got %v, want %v", h, 1.123)
		return
	}
}

func TestMapsStringBool(t *testing.T) {
	var m = tupi.Validate[map[string]bool](map[any]any{"bool": true})
	if m.HasErrors() {
		t.Error(m.Errors())
		return
	}

	if h, ok := m.Value()["bool"]; !ok {
		t.Errorf("TestMapString got %v, want %v", h, true)
		return
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
		return
	}

	mapUsers := m.Value()
	if len(mapUsers) != 6 {
		t.Errorf("len(TestMapString) got %v, want %v", len(mapUsers), len(data))
		return
	}
	if mapUsers["1"].Name != "etho" {
		t.Errorf("TestMapString['1'] got %v, want %v", mapUsers["1"], data["1"])
		return
	}
	if mapUsers["2"].Name != "etho2" {
		t.Errorf("TestMapString['2'] got %v, want %v", mapUsers["2"], data["2"])
		return
	}
	if mapUsers["3"].Name != "etho3" {
		t.Errorf("TestMapString['3'] got %v, want %v", mapUsers["3"], data["3"])
		return
	}
	if mapUsers["4"].Name != "etho4" {
		t.Errorf("TestMapString['4'] got %v, want %v", mapUsers["4"], data["4"])
		return
	}
	if mapUsers["5"].Name != "etho5" {
		t.Errorf("TestMapString['5'] got %v, want %v", mapUsers["5"], data["5"])
		return
	}
	if mapUsers["6"].Name != "etho6" {
		t.Errorf("TestMapString['6'] got %v, want %v", mapUsers["6"], data["6"])
		return
	}
}
