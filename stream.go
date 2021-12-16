package stream

type Stream[T any] struct {
	next func() []T
}

func Of[T any](elem ...T) *Stream[T] {
	next := func() []T {
		return elem
	}

	return &Stream[T]{next: next}
}

func (s *Stream[T]) Map(f func(it T) T) *Stream[T] {
	next := func() []T {
		list := s.next()
		if len(list) == 0 {
			return list
		}

		results := []T{}
		for i := range list {
			ptr := list[i]
			results = append(results, f(ptr))
		}

		return results
	}
	return &Stream[T]{next: next}
}

func (s *Stream[T]) Filter(f func(it T) bool) *Stream[T] {
	next := func() []T {
		list := s.next()
		if len(list) == 0 {
			return list
		}

		results := []T{}
		for i := range list {
			ptr := list[i]
			if f(ptr) {
				results = append(results, ptr)
			}
		}

		return results
	}
	return &Stream[T]{next: next}
}

func (s *Stream[T]) ForEach(f func(it T)) {
	list := s.next()
	for i := range list {
		f(list[i])
	}
}

func (s *Stream[T]) Collect() []T {
	return s.next()
}

func (s *Stream[T]) Any(f func(T) bool) bool {
	list := s.next()
	for i := range list {
		if f(list[i]) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) All(f func(T) bool) bool {
	list := s.next()
	for i := range list {
		if !f(list[i]) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) GroupBy(keySelector func(it T) any) map[any][]T {
	list := s.next()
	m := make(map[any][]T)

	for i := range list {
		ptr := list[i]
		key := keySelector(ptr)
		m[key] = append(m[key], ptr)
	}

	return m
}
