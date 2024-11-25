package gostd

func IfElse[T any](cond bool, t, f T) T {
	if cond {
		return t
	}
	return f
}
