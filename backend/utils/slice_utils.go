package slice_utils

func Map[T any, U any](slice []T, f func(T) U) []U {
	result := []U{}
	for _, e := range slice {
		result = append(result, f(e))
	}
	return result
}