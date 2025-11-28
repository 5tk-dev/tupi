package tupi

type Schema[T any] interface {
	Value() T
	Errors() []error
	HasErrors() bool
}

type schema[T any] struct {
	val    T
	errors []error
}

func (s *schema[T]) Value() T { return s.val }

func (s *schema[T]) Errors() []error { return s.errors }

func (s *schema[T]) HasErrors() bool { return len(s.errors) > 0 }
