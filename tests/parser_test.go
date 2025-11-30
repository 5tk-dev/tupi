package test

import (
	"testing"

	"5tk.dev/tupi"
)

type A struct {
	FieldA int
	FieldB int8
	FieldC int16
	FieldD int32
	FieldE int64
}

type B struct {
	FieldA int
	FieldB int8
	FieldC int16
	FieldD int32
	FieldE int64
}

type C struct {
	FieldA float32
	FieldB float64
}

type D struct {
	FieldA float32
	FieldB float64
}

type E struct {
	FieldA string
	FieldB string
}

type F struct {
	FieldA bool
	FieldB bool
}

type A1 struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

func BenchmarkParse_A(b *testing.B) {
	// tupi.Parse[A]()
	f := tupi.Fielder[*A]{}
	sch := f.Decode(map[string]string{
		"fieldA": "9223372036854775807",
		"fieldB": "127",
		"fieldC": "32767",
		"fieldD": "2147483647",
		"fieldE": "9223372036854775807",
	})
	if sch.HasErrors() {
		b.Error(sch.Errors())
	}
}

func BenchmarkParse_B(b *testing.B) {
	tupi.Parse[B]()
}

func BenchmarkParse_C(b *testing.B) {
	tupi.Parse[C]()
}

func BenchmarkParse_D(b *testing.B) {
	tupi.Parse[D]()
}

func BenchmarkParse_E(b *testing.B) {
	tupi.Parse[E]()
}

func BenchmarkParse_F(b *testing.B) {
	tupi.Parse[F]()
}

func BenchmarkParse_A1(b *testing.B) {
	tupi.Parse[A1]()
}

func BenchmarkDecodeWithoutParser(b *testing.B) {
	f := tupi.Fielder[*A]{}
	sch := f.Decode(map[string]string{
		"fieldA": "9223372036854775807",
		"fieldB": "127",
		"fieldC": "32767",
		"fieldD": "2147483647",
		"fieldE": "9223372036854775807",
	})
	if sch.HasErrors() {
		b.Error(sch.Errors())
		return
	}
	v := sch.Value()

	if v.FieldA != 9223372036854775807 {
		b.Errorf("got %q, want %d", v.FieldA, 9223372036854775807)
		return
	}
	if v.FieldB != 127 {
		b.Errorf("got %q, want %d", v.FieldB, 127)
		return
	}
	if v.FieldC != 32767 {
		b.Errorf("got %q, want %d", v.FieldC, 32767)
		return
	}
	if v.FieldD != 2147483647 {
		b.Errorf("got %q, want %d", v.FieldD, 2147483647)
		return
	}
	if v.FieldE != 9223372036854775807 {
		b.Errorf("got %q, want %d", v.FieldE, 9223372036854775807)
	}
}
