package test

import (
	"testing"

	"5tk.dev/tupi"
)

var (
	fielderA = &tupi.Fielder[*A]{}
	fielderC = &tupi.Fielder[*C]{}
	// fielderE  *tupi.Fielder[*E]
	// fielderF  *tupi.Fielder[*F]
	// fielderA1 *tupi.Fielder[*A1]
)

func BenchmarkParseAll(b *testing.B) {
	// fielderA = tupi.Parse[*A]()
	// fielderC = tupi.Parse[*C]()
	// fielderE = tupi.Parse[*E]()
	// fielderF = tupi.Parse[*F]()
	// fielderA1 = tupi.Parse[*A1]()
}

func BenchmarkDecode_A(b *testing.B) {
	for range 10 {
		sch := fielderA.Decode(map[string]string{
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
}

func BenchmarkDecode_C(b *testing.B) {
	sch := fielderC.Decode(map[string]string{
		"fieldA": "1.2147483",
		"fieldB": "1.92233720368547758",
	})
	if sch.HasErrors() {
		b.Error(sch.Errors())
	}
}

// func BenchmarkDecode_E(t *testing.T) {

// }

// func BenchmarkDecode_F(t *testing.T) {

// }

// func BenchmarkDecode_A1(t *testing.T) {

// }
