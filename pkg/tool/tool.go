package tool

func GetPointer[T any](value T) *T {
	return &value
}
