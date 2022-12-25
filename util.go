package heap

func swap[T SupportedTypes](a, b *T) {
	t := *a
	*a = *b
	*b = t
}
