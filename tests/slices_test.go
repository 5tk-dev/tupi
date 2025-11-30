package test

import (
	"testing"

	"5tk.dev/tupi"
)

func TestSliceInt(t *testing.T) {
	sch := tupi.Validate[[]int]([]any{1, "2", 3, 4, "5", "65"})
	// sch := tupi.Validate[[]int]([]any{1, 2, 3, 4, 5, 6})
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	if len(sch.Value()) != 6 {
		t.Errorf("TestMapString got %v, want %v", sch.Value(), []int{1, 2, 3, 4, 5, 6})
	}
}

func TestSliceFloat(t *testing.T) {
	sch := tupi.Validate[[]float32]([]any{1, "2", 3, 4, "5", "6.5678905"})
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	if len(sch.Value()) != 6 {
		t.Errorf("TestMapString got %v, want %v", sch.Value(), []int{1, 2, 3, 4, 5, 6})
	}
}

func TestSliceStruct(t *testing.T) {
	var data = []map[string]any{
		{"name": "etho", "age": 18},
		{"name": "etho", "age": 18},
		{"name": "etho", "age": 18},
		{"name": "etho", "age": 18},
	}
	sch := tupi.Validate[[]*UserTest](data)
	if sch.HasErrors() {
		t.Error(sch.Errors())
		return
	}
	if len(sch.Value()) != 4 {
		t.Errorf("TestMapString got %v, want %v", sch.Value(), []*UserTest{})
	}
}
