package test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"5tk.dev/tupi"
)

type SetRuleNowStruct struct {
	Now time.Time `validate:"now"`
}

func TestRuleNow(t *testing.T) {
	tupi.SetRule("now", &tupi.Rule{
		Validate: func(rv reflect.Value, rule string) bool {
			rv.Set(
				reflect.ValueOf(time.Now()),
			)
			return true
		},
	})
	sch := tupi.Validate[*SetRuleNowStruct](map[string]any{})
	if sch.HasErrors() {
		t.Error(sch.Errors())
	}
}

type SetRuleCpfStruct struct {
	Cpf string `validate:"cpf,minlength=11,maxlength=14"`
}

func TestRuleCpf(t *testing.T) {
	tupi.SetRule("cpf", &tupi.Rule{
		Validate: func(rv reflect.Value, rule string) bool {
			v := rv.Interface().(string)
			if len(v) != 9 {
				v = strings.ReplaceAll(v, ".", "")
				v = strings.ReplaceAll(v, "-", "")
			}
			rv.Set(reflect.ValueOf(v))
			return true
		},
	})
	f := tupi.Parse[*SetRuleCpfStruct]()

	cpf0 := map[string]string{"cpf": "000.000.000-00"}
	cpf1 := map[string]string{"cpf": "11111111111"}

	sch0 := f.Decode(cpf0)
	sch1 := f.Decode(cpf1)

	if sch0.HasErrors() {
		t.Error(sch0.Errors())
	}
	if sch1.HasErrors() {
		t.Error(sch1.Errors())
	}
	v0 := sch0.Value()
	if v0.Cpf != "00000000000" {
		t.Errorf("SetRuleCpfStruct got %v, want %v", v0.Cpf, "00000000000")

	}
	v1 := sch1.Value()
	if v1.Cpf != "11111111111" {
		t.Errorf("SetRuleCpfStruct got %v, want %v", v1.Cpf, "11111111111")

	}
}

type SetRuleEmailStruct struct {
	Email string `validate:"email,minlength=10,maxlength=64"`
}

func TestRuleEmail(t *testing.T) {
	tupi.SetRule("email", &tupi.Rule{
		Validate: func(rv reflect.Value, rule string) bool {
			v := rv.Interface().(string)
			strs := strings.Split(v, "@")
			if len(strs) != 2 {
				return false
			}
			if len(strs[0]) < 4 {
				return false
			}
			if len(strs[1]) < 4 {
				return false
			}
			return true
		},
	})

	f := tupi.Parse[*SetRuleEmailStruct]()
	sch0 := f.Decode(map[string]string{"email": "foobar@example.com"})
	if sch0.HasErrors() {
		t.Error(sch0.Errors())
		return
	}
	if v := sch0.Value(); v.Email != "foobar@example.com" {
		t.Errorf("SetRuleEmailStruct got %v, want %v", v.Email, "foobar@example.com")
		return
	}
	sch1 := f.Decode(map[string]string{"email": "foo@example.com"})
	if !sch1.HasErrors() {
		v := sch1.Value()
		t.Errorf("SetRuleEmailStruct got %v, want %v", v.Email, "nil")
		return
	}
	sch2 := f.Decode(map[string]string{"email": "foobar@com"})
	if !sch2.HasErrors() {
		v := sch2.Value()
		t.Errorf("SetRuleEmailStruct got %v, want %v", v.Email, "nil")
		return
	}

	sch10 := f.Decode(map[string]string{"email": "foo.com"})
	if !sch10.HasErrors() {
		v := sch10.Value()
		t.Errorf("SetRuleEmailStruct got %v, want %v", v.Email, "nil")
		return
	}

}
