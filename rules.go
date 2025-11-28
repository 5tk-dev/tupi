package tupi

import (
	"fmt"
	"reflect"
	"strings"
)

type Rule struct {
	Name           string //
	Value          string //
	Message        string // ex.: {field} require a value >= 18
	Validate       func(rv reflect.Value, rule string) bool
	BeforeSetValue bool // exec BEFORE set value
}

func (r *Rule) ToMap() map[string]any {
	return map[string]any{
		"name":         r.Name,
		"value":        r.Value,
		"errorMessage": r.Message,
	}
}

var defaultRules = map[string]bool{
	"":         true,
	"required": true,
}

var (
	rules = map[string]*Rule{
		"required": {
			Name:     "required",
			Message:  "{field} is required",
			Validate: req,
		},
		"escape": {
			Name:     "escape",
			Message:  "{field} do not be replaced",
			Validate: escape,
		},
		"min": {
			Name:     "min",
			Message:  "{field} requires a value >= {value}",
			Validate: min,
		},
		"max": {
			Name:     "max",
			Message:  "{field} requires a value <= {value}",
			Validate: max,
		},
		"minlen": {
			Name:     "minlen",
			Message:  "{field} requires a length value >= {value}",
			Validate: minLen,
		},
		"maxlen": {
			Name:     "maxlen",
			Message:  "{field} requires a length value <= {value}",
			Validate: maxLen,
		},
	}
)

// ex:
//
//		tupi.SetRules("min",&tupi.Rule{Message:"min value: {value}"}) // '{value}' be replace by "Rule.Value"
//		tupi.SetRules("max",&tupi.Rule{Message:"max value: {value}"}) // '{value}' be replace by "Rule.Value"
//		tupi.SetRules("format",&tupi.Rule{Validate:func(value any) bool {...}})
//
//		type User struct {
//			Age int `tupi:"min=18"`
//	 }
func SetRule(field string, rule *Rule) {
	rule.Name = strings.ToLower(field)
	if _, ok := defaultRules[rule.Name]; ok {
		panic(fmt.Errorf("%q ia a invalid tag rule", rule.Name))
	}
	rules[rule.Name] = rule
}

func GetRule(rule string) *Rule { return rules[rule] }
