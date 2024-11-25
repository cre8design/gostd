package gostd

import "iter"

func Map[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for a := range seq {
			if !yield(f(a)) {
				return
			}
		}
	}
}

func Filter[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for a := range seq {
			if f(a) {
				if !yield(a) {
					return
				}
			}
		}
	}
}

func ForEach[T any](seq iter.Seq[T], f func(T)) {
	for a := range seq {
		f(a)
	}
}

func Every[T any](seq iter.Seq[T], f func(T) bool) bool {
	for a := range seq {
		if !f(a) {
			return false
		}
	}
	return true
}

func Some[T any](seq iter.Seq[T], f func(T) bool) bool {
	for a := range seq {
		if f(a) {
			return true
		}
	}
	return false
}

func Find[T any](seq iter.Seq[T], f func(T) bool) (*T, bool) {
	for a := range seq {
		if f(a) {
			return &a, true
		}
	}
	return nil, false
}

func IndexOf[T any](seq iter.Seq[T], f func(T) bool) int {
	i := 0
	for a := range seq {
		if f(a) {
			return i
		}
		i++
	}
	return -1
}
