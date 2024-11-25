package gostd

func KeysOf[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func ValuesOf[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func KeyByValue[K comparable, V comparable](m map[K]V, v V) (*K, bool) {
	for k, vv := range m {
		if v == vv {
			return &k, true
		}
	}
	return nil, false
}
