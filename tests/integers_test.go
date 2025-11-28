package test

import (
	"testing"

	"5tk.dev/tupi"
)

var (
	fielderInt   = tupi.ParseSchema[int]() // alias int64
	fielderInt8  = tupi.ParseSchema[int8]()
	fielderInt16 = tupi.ParseSchema[int16]()
	fielderInt32 = tupi.ParseSchema[int32]()
)

// int.Decode(int) -> !Schema.hasError()
func TestIntegersWithInts(t *testing.T) {
	dataInt8 := 127
	dataInt16 := 32767
	dataInt32 := 2147483647
	dataInt64 := 9223372036854775807

	/* INT8 */
	schInt8 := fielderInt8.Decode(dataInt8)
	if schInt8.HasErrors() {
		t.Error(schInt8.Errors())
	}
	vInt8 := schInt8.Value()
	if vInt8 != 127 {
		t.Errorf("got %v, want %v", vInt8, 127)
	}

	/* INT16 */
	schInt16 := fielderInt16.Decode(dataInt16)
	if schInt16.HasErrors() {
		t.Error(schInt16.Errors())
	}
	vInt16 := schInt16.Value()
	if vInt16 != 32767 {
		t.Errorf("got %v, want %v", vInt16, 32767)
	}

	/* INT32 */
	schInt32 := fielderInt32.Decode(dataInt32)
	if schInt32.HasErrors() {
		t.Error(schInt32.Errors())
	}
	vInt32 := schInt32.Value()
	if vInt32 != 2147483647 {
		t.Errorf("got %v, want %v", vInt32, 2147483647)
	}

	/* INT OR INT64 */

	schInt := fielderInt.Decode(dataInt64)
	if schInt.HasErrors() {

		t.Error(schInt.Errors())
	}
	vInt := schInt.Value()
	if vInt != 9223372036854775807 {
		t.Errorf("got %v, want %v", vInt, dataInt64)
	}
}

// int.Decode("int") -> !Schema.hasError()
func TestIntegersWithStringInts(t *testing.T) {
	dataInt8 := "127"
	dataInt16 := "32767"
	dataInt32 := "2147483647"
	dataInt64 := "9223372036854775807"

	/* INT8 */
	schInt8 := fielderInt8.Decode(dataInt8)
	if schInt8.HasErrors() {
		t.Error(schInt8.Errors())
	}
	vInt8 := schInt8.Value()
	if vInt8 != 127 {
		t.Errorf("got %v, want %v", vInt8, 127)
	}

	/* INT16 */
	schInt16 := fielderInt16.Decode(dataInt16)
	if schInt16.HasErrors() {
		t.Error(schInt16.Errors())
	}
	vInt16 := schInt16.Value()
	if vInt16 != 32767 {
		t.Errorf("got %v, want %v", vInt16, 32767)
	}

	/* INT32 */
	schInt32 := fielderInt32.Decode(dataInt32)
	if schInt32.HasErrors() {
		t.Error(schInt32.Errors())
	}
	vInt32 := schInt32.Value()
	if vInt32 != 2147483647 {
		t.Errorf("got %v, want %v", vInt32, 2147483647)
	}

	/* INT OR INT64 */

	schInt := fielderInt.Decode(dataInt64)
	if schInt.HasErrors() {

		t.Error(schInt.Errors())
	}
	vInt := schInt.Value()
	if vInt != 9223372036854775807 {
		t.Errorf("got %v, want %v", vInt, dataInt64)
	}
}

// // int.Decode(float) -> !Schema.hasError()
func TestIntegersWithFloats(t *testing.T) {
	dataInt8 := 127.0
	dataInt16 := 32767.0
	dataInt32 := 2147483647.0
	dataInt64 := 8223372036854775807.1 // isso aq é doidera

	/* INT8 */
	schInt8 := fielderInt8.Decode(dataInt8)
	if schInt8.HasErrors() {
		t.Error(schInt8.Errors())
	}
	vInt8 := schInt8.Value()
	if vInt8 != 127 {
		t.Errorf("got %v, want %v", vInt8, 127)
	}

	/* INT16 */
	schInt16 := fielderInt16.Decode(dataInt16)
	if schInt16.HasErrors() {
		t.Error(schInt16.Errors())
	}
	vInt16 := schInt16.Value()
	if vInt16 != 32767 {
		t.Errorf("got %v, want %v", vInt16, 32767)
	}

	/* INT32 */
	schInt32 := fielderInt32.Decode(dataInt32)
	if schInt32.HasErrors() {
		t.Error(schInt32.Errors())
	}
	vInt32 := schInt32.Value()
	if vInt32 != 2147483647 {
		t.Errorf("got %v, want %v", vInt32, 2147483647)
	}

	/* INT OR INT64 */

	schInt := fielderInt.Decode(dataInt64)
	if schInt.HasErrors() {

		t.Error(schInt.Errors())
	}
	vInt := schInt.Value()
	if vInt != 8223372036854775808 {
		t.Errorf("got %v, want %d", vInt, 8223372036854775808)
	}
}

// int.Decode("string") -> Schema.hasError()
func TestIntegersWithStringValues(t *testing.T) {
	sch8 := fielderInt8.Decode("fdagaj156171")
	if !sch8.HasErrors() {
		t.Errorf("got %q, want %q", sch8.Value(), "error")
	}
	sch16 := fielderInt16.Decode("a")
	if !sch16.HasErrors() {
		t.Errorf("got %q, want %q", sch16.Value(), "error")
	}
	sch32 := fielderInt32.Decode("cdfc")
	if !sch32.HasErrors() {
		t.Errorf("got %q, want %q", sch32.Value(), "error")
	}
	sch64 := fielderInt.Decode("\t\t")
	if !sch64.HasErrors() {
		t.Errorf("got %q, want %q", sch64.Value(), "error")
	}
}

// // int.Decode(false) -> Schema.hasError()
func TestIntegersWithBooleans0(t *testing.T) {
	sch8 := fielderInt8.Decode(false)
	if !sch8.HasErrors() {
		t.Errorf("got %q, want %q", sch8.Value(), "error")
	}
	sch16 := fielderInt16.Decode(false)
	if !sch16.HasErrors() {
		t.Errorf("got %q, want %q", sch16.Value(), "error")
	}
	sch32 := fielderInt32.Decode(false)
	if !sch32.HasErrors() {
		t.Errorf("got %q, want %q", sch32.Value(), "error")
	}
	sch64 := fielderInt.Decode(false)
	if !sch64.HasErrors() {
		t.Errorf("got %q, want %q", sch64.Value(), "error")
	}
}

// // int.Decode(true) -> Schema.hasError()
func TestIntegersWithBooleans1(t *testing.T) {
	sch8 := fielderInt8.Decode(true)
	if !sch8.HasErrors() {
		t.Errorf("got %q, want %q", sch8.Value(), "error")
	}
	sch16 := fielderInt16.Decode(true)
	if !sch16.HasErrors() {
		t.Errorf("got %q, want %q", sch16.Value(), "error")
	}
	sch32 := fielderInt32.Decode(true)
	if !sch32.HasErrors() {
		t.Errorf("got %q, want %q", sch32.Value(), "error")
	}
	sch64 := fielderInt.Decode(true)
	if !sch64.HasErrors() {
		t.Errorf("got %q, want %q", sch64.Value(), "error")
	}
}

// // int.Decode(struct) -> Schema.hasError()
func TestIntegersWithStructs(t *testing.T) {
	s := struct{}{}
	sch8 := fielderInt8.Decode(s)
	if !sch8.HasErrors() {
		t.Errorf("got %q, want %q", sch8.Value(), "error")
	}
	sch16 := fielderInt16.Decode(s)
	if !sch16.HasErrors() {
		t.Errorf("got %q, want %q", sch16.Value(), "error")
	}
	sch32 := fielderInt32.Decode(s)
	if !sch32.HasErrors() {
		t.Errorf("got %q, want %q", sch32.Value(), "error")
	}
	sch64 := fielderInt.Decode(s)
	if !sch64.HasErrors() {
		t.Errorf("got %q, want %q", sch64.Value(), "error")
	}
}

// // int.Decode(map) -> Schema.hasError()
func TestIntegersWithMaps(t *testing.T) {
	m := map[string]any{}
	sch8 := fielderInt8.Decode(m)
	if !sch8.HasErrors() {
		t.Errorf("got %q, want %q", sch8.Value(), "error")
	}
	sch16 := fielderInt16.Decode(m)
	if !sch16.HasErrors() {
		t.Errorf("got %q, want %q", sch16.Value(), "error")
	}
	sch32 := fielderInt32.Decode(m)
	if !sch32.HasErrors() {
		t.Errorf("got %q, want %q", sch32.Value(), "error")
	}
	sch64 := fielderInt.Decode(m)
	if !sch64.HasErrors() {
		t.Errorf("got %q, want %q", sch64.Value(), "error")
	}
}

// // int.Decode(func) -> Schema.hasError()
func TestIntegersWithFuncs(t *testing.T) {
	f := func() {}
	sch8 := fielderInt8.Decode(f)
	if !sch8.HasErrors() {
		t.Errorf("got %q, want %q", sch8.Value(), "error")
	}
	sch16 := fielderInt16.Decode(f)
	if !sch16.HasErrors() {
		t.Errorf("got %q, want %q", sch16.Value(), "error")
	}
	sch32 := fielderInt32.Decode(f)
	if !sch32.HasErrors() {
		t.Errorf("got %q, want %q", sch32.Value(), "error")
	}
	sch64 := fielderInt.Decode(f)
	if !sch64.HasErrors() {
		t.Errorf("got %q, want %q", sch64.Value(), "error")
	}
}
