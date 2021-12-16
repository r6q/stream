package sequence

type Seq[T any] struct {
	next func() *T
}

func Of[T any](elem ...T) *Seq[T] {
	next := func() *T {
		if len(elem) < 1 {
			return nil
		}
		first := &elem[0]
		elem = elem[1:]
		return first
	}

	return &Seq[T]{next: next}
}

func (s *Seq[T]) Map(f func(it T) T) *Seq[T] {
	next := func() *T {
		ptr := s.next()
		if ptr == nil {
			return nil
		}
		*ptr = f(*ptr)
		return ptr
	}
	return &Seq[T]{next: next}
}

func (s *Seq[T]) Filter(f func(it T) bool) *Seq[T] {
	next := func() *T {
		for {
			ptr := s.next()
			if ptr == nil {
				return nil
			}
			if f(*ptr) {
				return ptr
			}
		}
	}
	return &Seq[T]{next: next}
}

func (s *Seq[T]) ForEach(f func(it T)) {
	for ptr := s.next(); ptr != nil; ptr = s.next() {
		f(*ptr)
	}
}

func (s *Seq[T]) Collect() []T {
	slice := []T{}
	for ptr := s.next(); ptr != nil; ptr = s.next() {
		slice = append(slice, *ptr)
	}
	return slice
}

func (s *Seq[T]) Any(f func(it T) bool) bool {
	for ptr := s.next(); ptr != nil; ptr = s.next() {
		if f(*ptr) {
			return true
		}
	}
	return false
}
