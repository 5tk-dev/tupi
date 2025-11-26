package tupi

type Schema interface {
	Value() any
	Errors() []error
	HasErrors() bool
}

type schema struct {
	val    any
	errors []error
}

func (s *schema) Value() any { return s.val }

func (s *schema) Errors() []error { return s.errors }

func (s *schema) HasErrors() bool { return len(s.errors) > 0 }
