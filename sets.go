package xtreamcodes

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(e T) Set[T] {
	s[e] = struct{}{}
	return s
}

func (s Set[T]) Remove(e T) Set[T] {
	delete(s, e)
	return s
}

func (s Set[T]) Has(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) ToArray() []T {
	v := make([]T, 0)
	for k, _ := range s {
		v = append(v, k)
	}
	return v
}

func From[T comparable](elms ...T) Set[T] {
	s := make(Set[T], len(elms))
	for _, e := range elms {
		s.Add(e)
	}
	return s
}
