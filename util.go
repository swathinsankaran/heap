package heap

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
