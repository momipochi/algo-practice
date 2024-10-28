package utils

import "testing"

func IsEqual2DArray(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
func IsEqual2DArrayString(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func FlattenArray[T any](collection [][]T) []T {
	res := []T{}
	for i := 0; i < len(collection); i++ {
		for j := 0; j < len(collection[i]); j++ {
			res = append(res, collection[i][j])
		}
	}
	return res
}

func IsEqualArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if (a)[i] != (b)[i] {
			return false
		}
	}
	return true
}
func IsEqualArrayString(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if (a)[i] != (b)[i] {
			return false
		}
	}
	return true
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func DeafultTestOutput[V any](t *testing.T, expected V, actual V) {
	t.Errorf("Wrong answer. \nExpected <%v>\nReceived <%v>", expected, actual)
}
