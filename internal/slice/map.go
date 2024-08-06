package slice

func Map[T1, T2 any](list []T1, yield func(T1) T2) []T2 {
	newList := make([]T2, 0, len(list))
	for _, t := range list {
		newList = append(newList, yield(t))
	}
	return newList
}
