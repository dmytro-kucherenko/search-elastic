package utils

import (
	"fmt"
	"slices"
)

func Unique[T comparable](values []T) []T {
	result := make([]T, 0)

	for _, value := range values {
		if !slices.Contains(result, value) {
			result = append(result, value)
		}
	}

	return result
}

func Union[T comparable](values ...[]T) []T {
	result := make([]T, 0)
	if len(values) == 0 {
		return result
	}

	for _, value := range values {
		result = append(result, value...)
	}

	return Unique(result)
}

func Intersection[T comparable](values ...[]T) []T {
	result := make([]T, 0)
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		stepResult := make([]T, 0, len(result))
		for _, value := range result {
			if slices.Contains(values[i], value) {
				stepResult = append(stepResult, value)
			}
		}

		result = stepResult
	}

	fmt.Println(values, result)

	return result
}
