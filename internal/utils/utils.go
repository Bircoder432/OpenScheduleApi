package utils

func ToNewSlice[S any, Z any](slice []S, f func(S) (Z, bool)) []Z {
	newSlice := make([]Z, 0, len(slice))
	for _, v := range slice {
		if newValue, ok := f(v); ok {
			newSlice = append(newSlice, newValue)
		}
	}
	return newSlice
}
